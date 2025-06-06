// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/frontend/v1/frontend.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package v1

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/dskit/httpgrpc"
	"github.com/grafana/dskit/services"
	"github.com/grafana/dskit/tenant"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/grafana/mimir/pkg/frontend/v1/frontendv1pb"
	"github.com/grafana/mimir/pkg/querier/stats"
	"github.com/grafana/mimir/pkg/scheduler/queue"
	"github.com/grafana/mimir/pkg/util"
	"github.com/grafana/mimir/pkg/util/httpgrpcutil"
	"github.com/grafana/mimir/pkg/util/validation"
)

var (
	errTooManyRequest = httpgrpc.Errorf(http.StatusTooManyRequests, "too many outstanding requests")
)

var tracer = otel.Tracer("pkg/frontend/v1")

// Config for a Frontend.
type Config struct {
	MaxOutstandingPerTenant int           `yaml:"max_outstanding_per_tenant" category:"advanced"`
	QuerierForgetDelay      time.Duration `yaml:"querier_forget_delay" category:"experimental"`
}

// RegisterFlags adds the flags required to config this to the given FlagSet.
func (cfg *Config) RegisterFlags(f *flag.FlagSet) {
	f.IntVar(&cfg.MaxOutstandingPerTenant, "querier.max-outstanding-requests-per-tenant", 100, "Maximum number of outstanding requests per tenant per frontend; requests beyond this error with HTTP 429.")
	f.DurationVar(&cfg.QuerierForgetDelay, "query-frontend.querier-forget-delay", 0, "If a querier disconnects without sending notification about graceful shutdown, the query-frontend will keep the querier in the tenant's shard until the forget delay has passed. This feature is useful to reduce the blast radius when shuffle-sharding is enabled.")
}

type Limits interface {
	// Returns max queriers to use per tenant, or 0 if shuffle sharding is disabled.
	MaxQueriersPerUser(user string) int
}

// Frontend queues HTTP requests, dispatches them to backends, and handles retries
// for requests which failed.
type Frontend struct {
	services.Service

	cfg    Config
	log    log.Logger
	limits Limits

	requestQueue *queue.RequestQueue
	activeUsers  *util.ActiveUsersCleanupService

	// Subservices manager.
	subservices        *services.Manager
	subservicesWatcher *services.FailureWatcher

	// Metrics.
	queueLength       *prometheus.GaugeVec
	discardedRequests *prometheus.CounterVec
	numClients        prometheus.GaugeFunc
	queueDuration     prometheus.Histogram
}

type request struct {
	enqueueTime time.Time
	queueSpan   trace.Span
	originalCtx context.Context

	request  *httpgrpc.HTTPRequest
	err      chan error
	response chan *httpgrpc.HTTPResponse
}

// New creates a new frontend. Frontend implements service, and must be started and stopped.
func New(cfg Config, limits Limits, log log.Logger, registerer prometheus.Registerer) (*Frontend, error) {
	var err error

	f := &Frontend{
		cfg:                cfg,
		log:                log,
		limits:             limits,
		subservicesWatcher: services.NewFailureWatcher(),
		queueLength: promauto.With(registerer).NewGaugeVec(prometheus.GaugeOpts{
			Name: "cortex_query_frontend_queue_length",
			Help: "Number of queries in the queue.",
		}, []string{"user"}),
		discardedRequests: promauto.With(registerer).NewCounterVec(prometheus.CounterOpts{
			Name: "cortex_query_frontend_discarded_requests_total",
			Help: "Total number of query requests discarded.",
		}, []string{"user"}),
		queueDuration: promauto.With(registerer).NewHistogram(prometheus.HistogramOpts{
			Name:    "cortex_query_frontend_queue_duration_seconds",
			Help:    "Time spent by requests in queue before getting picked up by a querier.",
			Buckets: prometheus.DefBuckets,
		}),
	}

	enqueueDuration := promauto.With(registerer).NewHistogram(prometheus.HistogramOpts{
		Name: "cortex_query_frontend_enqueue_duration_seconds",
		Help: "Time spent by requests waiting to join the queue or be rejected.",
	})
	querierInflightRequests := promauto.With(registerer).NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "cortex_query_frontend_querier_inflight_requests",
			Help:       "Number of inflight requests being processed on all querier-scheduler connections. Quantile buckets keep track of inflight requests over the last 60s.",
			Objectives: map[float64]float64{0.5: 0.05, 0.75: 0.02, 0.8: 0.02, 0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			MaxAge:     time.Minute,
			AgeBuckets: 6,
		},
		[]string{"query_component"},
	)
	f.requestQueue, err = queue.NewRequestQueue(
		log,
		cfg.MaxOutstandingPerTenant,
		cfg.QuerierForgetDelay,
		f.queueLength,
		f.discardedRequests,
		enqueueDuration,
		querierInflightRequests,
	)
	if err != nil {
		return nil, err
	}

	f.activeUsers = util.NewActiveUsersCleanupWithDefaultValues(f.cleanupInactiveUserMetrics)

	f.subservices, err = services.NewManager(f.requestQueue, f.activeUsers)
	if err != nil {
		return nil, err
	}

	f.numClients = promauto.With(registerer).NewGaugeFunc(prometheus.GaugeOpts{
		Name: "cortex_query_frontend_connected_clients",
		Help: "Number of worker clients currently connected to the frontend.",
	}, f.requestQueue.GetConnectedQuerierWorkersMetric)

	f.Service = services.NewBasicService(f.starting, f.running, f.stopping)
	return f, nil
}

func (f *Frontend) starting(ctx context.Context) error {
	f.subservicesWatcher.WatchManager(f.subservices)

	if err := services.StartManagerAndAwaitHealthy(ctx, f.subservices); err != nil {
		return errors.Wrap(err, "unable to start frontend subservices")
	}

	return nil
}

func (f *Frontend) running(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-f.subservicesWatcher.Chan():
			return errors.Wrap(err, "frontend subservice failed")
		}
	}
}

func (f *Frontend) stopping(_ error) error {
	// This will also stop the requests queue, which stop accepting new requests and errors out any pending requests.
	return services.StopManagerAndAwaitStopped(context.Background(), f.subservices)
}

func (f *Frontend) cleanupInactiveUserMetrics(user string) {
	f.queueLength.DeleteLabelValues(user)
	f.discardedRequests.DeleteLabelValues(user)
}

// RoundTripGRPC round trips a proto (instead of an HTTP request).
func (f *Frontend) RoundTripGRPC(ctx context.Context, req *httpgrpc.HTTPRequest) (*httpgrpc.HTTPResponse, io.ReadCloser, error) {
	// Propagate trace context in gRPC too - this will be ignored if using HTTP.
	otel.GetTextMapPropagator().Inject(ctx, (*httpgrpcutil.HttpgrpcHeadersCarrier)(req))

	request := request{
		request:     req,
		originalCtx: ctx,

		// Buffer of 1 to ensure response can be written by the server side
		// of the Process stream, even if this goroutine goes away due to
		// client context cancellation.
		err:      make(chan error, 1),
		response: make(chan *httpgrpc.HTTPResponse, 1),
	}

	if err := f.queueRequest(ctx, &request); err != nil {
		return nil, nil, err
	}

	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()

	case resp := <-request.response:
		return resp, nil, nil

	case err := <-request.err:
		return nil, nil, err
	}
}

// Process allows backends to pull requests from the frontend.
func (f *Frontend) Process(server frontendv1pb.Frontend_ProcessServer) error {
	querierID, err := getQuerierID(server)
	if err != nil {
		return err
	}

	querierWorkerConn := queue.NewUnregisteredQuerierWorkerConn(server.Context(), querierID)
	err = f.requestQueue.AwaitRegisterQuerierWorkerConn(querierWorkerConn)
	if err != nil {
		return err
	}
	defer f.requestQueue.SubmitUnregisterQuerierWorkerConn(querierWorkerConn)

	lastTenantIdx := queue.FirstTenant()

	for {
		dequeueReq := queue.NewQuerierWorkerDequeueRequest(querierWorkerConn, lastTenantIdx)
		reqWrapper, idx, err := f.requestQueue.AwaitRequestForQuerier(dequeueReq)
		if err != nil {
			return err
		}
		lastTenantIdx = idx

		req := reqWrapper.(*request)

		queueTime := time.Since(req.enqueueTime)
		f.queueDuration.Observe(queueTime.Seconds())
		req.queueSpan.End()

		/*
		  We want to dequeue the next unexpired request from the chosen tenant queue.
		  The chance of choosing a particular tenant for dequeuing is (1/active_tenants).
		  This is problematic under load, especially with other middleware enabled such as
		  querier.split-by-interval, where one request may fan out into many.
		  If expired requests aren't exhausted before checking another tenant, it would take
		  n_active_tenants * n_expired_requests_at_front_of_queue requests being processed
		  before an active request was handled for the tenant in question.
		  If this tenant meanwhile continued to queue requests,
		  it's possible that its own queue would perpetually contain only expired requests.
		*/
		if req.originalCtx.Err() != nil {
			lastTenantIdx = lastTenantIdx.ReuseLastTenant()
			continue
		}

		// Handle the stream sending & receiving on a goroutine so we can
		// monitor the contexts in a select and cancel things appropriately.
		resps := make(chan *frontendv1pb.ClientToFrontend, 1)
		errs := make(chan error, 1)
		go func() {
			err = server.Send(&frontendv1pb.FrontendToClient{
				Type:           frontendv1pb.HTTP_REQUEST,
				HttpRequest:    req.request,
				StatsEnabled:   stats.IsEnabled(req.originalCtx),
				QueueTimeNanos: queueTime.Nanoseconds(),
			})
			if err != nil {
				errs <- err
				return
			}

			resp, err := server.Recv()
			if err != nil {
				errs <- err
				return
			}

			resps <- resp
		}()

		select {
		// If the upstream request is cancelled, we need to cancel the
		// downstream req.  Only way we can do that is to close the stream.
		// The worker client is expecting this semantics.
		case <-req.originalCtx.Done():
			return req.originalCtx.Err()

		// Is there was an error handling this request due to network IO,
		// then error out this upstream request _and_ stream.
		case err := <-errs:
			req.err <- err
			return err

		// Happy path: merge the stats and propagate the response.
		case resp := <-resps:
			if stats.ShouldTrackHTTPGRPCResponse(resp.HttpResponse) {
				stats := stats.FromContext(req.originalCtx)
				stats.Merge(resp.Stats) // Safe if stats is nil.
			}

			req.response <- resp.HttpResponse
		}
	}
}

func (f *Frontend) NotifyClientShutdown(ctx context.Context, req *frontendv1pb.NotifyClientShutdownRequest) (*frontendv1pb.NotifyClientShutdownResponse, error) {
	level.Info(f.log).Log("msg", "received shutdown notification from querier", "querier", req.GetClientID())
	f.requestQueue.SubmitNotifyQuerierShutdown(ctx, req.GetClientID())

	return &frontendv1pb.NotifyClientShutdownResponse{}, nil
}

func getQuerierID(server frontendv1pb.Frontend_ProcessServer) (string, error) {
	err := server.Send(&frontendv1pb.FrontendToClient{
		Type: frontendv1pb.GET_ID,
		// Old queriers don't support GET_ID, and will try to use the request.
		// To avoid confusing them, include dummy request.
		HttpRequest: &httpgrpc.HTTPRequest{
			Method: "GET",
			Url:    "/invalid_request_sent_by_frontend",
		},
	})

	if err != nil {
		return "", err
	}

	resp, err := server.Recv()

	// Old queriers will return empty string, which is fine. All old queriers will be
	// treated as single querier with lot of connections.
	// (Note: if resp is nil, GetClientID() returns "")
	return resp.GetClientID(), err
}

func (f *Frontend) queueRequest(ctx context.Context, req *request) error {
	tenantIDs, err := tenant.TenantIDs(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	req.enqueueTime = now
	_, req.queueSpan = tracer.Start(ctx, "queued")

	// aggregate the max queriers limit in the case of a multi tenant query
	maxQueriers := validation.SmallestPositiveNonZeroIntPerTenant(tenantIDs, f.limits.MaxQueriersPerUser)

	joinedTenantID := tenant.JoinTenantIDs(tenantIDs)
	f.activeUsers.UpdateUserTimestamp(joinedTenantID, now)

	err = f.requestQueue.SubmitRequestToEnqueue(joinedTenantID, req, maxQueriers, nil)
	if errors.Is(err, queue.ErrTooManyRequests) {
		return errTooManyRequest
	}
	return err
}

// CheckReady determines if the query frontend is ready.  Function parameters/return
// chosen to match the same method in the ingester
func (f *Frontend) CheckReady(_ context.Context) error {
	// if we have more than one querier connected we will consider ourselves ready
	connectedClients := f.requestQueue.GetConnectedQuerierWorkersMetric()
	if connectedClients > 0 {
		return nil
	}

	msg := fmt.Sprintf("not ready: number of queriers connected to query-frontend is %d", int64(connectedClients))
	level.Info(f.log).Log("msg", msg)
	return errors.New(msg)
}
