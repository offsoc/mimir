// SPDX-License-Identifier: AGPL-3.0-only

package alertmanager

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-kit/log"
	"github.com/grafana/alerting/definition"
	"github.com/stretchr/testify/require"

	"github.com/grafana/mimir/pkg/alertmanager/alertspb"
)

const grafanaConfigWithDuplicateReceiverName = `{
  "template_files": {},
  "alertmanager_config": {
    "route": {
      "receiver": "test-receiver8",
      "group_by": [
        "grafana_folder",
        "alertname"
      ]
    },
    "templates": null,
    "receivers": [
      {"name": "test-receiver1","grafana_managed_receiver_configs": []},
      {
        "name": "test-receiver8",
        "grafana_managed_receiver_configs": [
          {
            "uid": "dde6ntuob69dtf",
            "name": "test-receiver",
            "type": "webhook",
            "disableResolveMessage": false,
            "settings": {
              "url": "http://localhost:8080",
              "username": "test"
            },
            "secureSettings": {
              "password": "test"
            }
          }
        ]
      },
      {"name": "test-receiver2","grafana_managed_receiver_configs": []},
      {"name": "test-receiver3","grafana_managed_receiver_configs": []},
      {"name": "test-receiver4","grafana_managed_receiver_configs": []},
      {"name": "test-receiver5","grafana_managed_receiver_configs": []},
      {"name": "test-receiver6","grafana_managed_receiver_configs": []},
      {"name": "test-receiver7","grafana_managed_receiver_configs": []},
      {
        "name": "test-receiver8",
        "grafana_managed_receiver_configs": [
          {
            "uid": "dde7ntuob69dtf",
            "name": "test-receiver",
            "type": "webhook",
            "disableResolveMessage": false,
            "settings": {
              "url": "http://localhost:8080",
              "username": "test"
            },
            "secureSettings": {
              "password": "test"
            }
          }
        ]
      },
	  {"name": "test-receiver9","grafana_managed_receiver_configs": []}
    ]
  }
}`

func TestCreateUsableGrafanaConfig(t *testing.T) {
	tests := []struct {
		name          string
		grafanaConfig alertspb.GrafanaAlertConfigDesc
		mimirConfig   string
		expErr        string
	}{
		{
			"empty grafana config",
			alertspb.GrafanaAlertConfigDesc{
				ExternalUrl:   "http://test:3000",
				RawConfig:     "",
				StaticHeaders: map[string]string{"test": "test"},
				SmtpFrom:      "test-instance@grafana.com",
			},
			simpleConfigOne,
			"failed to unmarshal Grafana Alertmanager configuration: unexpected end of JSON input",
		},
		{
			"invalid grafana config",
			alertspb.GrafanaAlertConfigDesc{
				ExternalUrl:   "http://test:3000",
				RawConfig:     "invalid",
				StaticHeaders: map[string]string{"test": "test"},
				SmtpFrom:      "test-instance@grafana.com",
			},
			simpleConfigOne,
			"failed to unmarshal Grafana Alertmanager configuration: invalid character 'i' looking for beginning of value",
		},
		{
			"no mimir config",
			alertspb.GrafanaAlertConfigDesc{
				ExternalUrl:   "http://test:3000",
				RawConfig:     grafanaConfig,
				StaticHeaders: map[string]string{"test": "test"},
				SmtpFrom:      "test-instance@grafana.com",
			},
			"",
			"",
		},
		{
			"duplicate grafana receiver name config",
			alertspb.GrafanaAlertConfigDesc{
				ExternalUrl:   "http://test:3000",
				RawConfig:     grafanaConfigWithDuplicateReceiverName,
				StaticHeaders: map[string]string{"test": "test"},
				SmtpFrom:      "test-instance@grafana.com",
			},
			"",
			"",
		},
		{
			"non-empty mimir config",
			alertspb.GrafanaAlertConfigDesc{
				ExternalUrl:   "http://test:3000",
				RawConfig:     grafanaConfig,
				StaticHeaders: map[string]string{"test": "test"},
				SmtpFrom:      "test-instance@grafana.com",
			},
			simpleConfigOne,
			"",
		},
		{
			"non-empty mimir config, empty SMTP from address",
			alertspb.GrafanaAlertConfigDesc{
				ExternalUrl:   "http://test:3000",
				RawConfig:     grafanaConfig,
				StaticHeaders: map[string]string{"test": "test"},
			},
			simpleConfigOne,
			"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			am := MultitenantAlertmanager{logger: log.NewNopLogger()}
			cfg, err := createUsableGrafanaConfig(am.logger, test.grafanaConfig, test.mimirConfig)
			if test.expErr != "" {
				require.Error(t, err)
				require.Equal(t, test.expErr, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.grafanaConfig.StaticHeaders, cfg.staticHeaders)
			require.Equal(t, test.grafanaConfig.User, cfg.User)
			require.Equal(t, test.grafanaConfig.ExternalUrl, cfg.tmplExternalURL.String())
			require.True(t, cfg.usingGrafanaConfig)

			if test.mimirConfig != "" {
				// The resulting config should contain Mimir's globals with Grafana's "from" address.
				mCfg, err := definition.LoadCompat([]byte(test.mimirConfig))
				require.NoError(t, err)

				var gCfg GrafanaAlertmanagerConfig
				require.NoError(t, json.Unmarshal([]byte(test.grafanaConfig.RawConfig), &gCfg))

				gCfg.AlertmanagerConfig.Global = mCfg.Global
				if test.grafanaConfig.SmtpFrom != "" {
					gCfg.AlertmanagerConfig.Global.SMTPFrom = test.grafanaConfig.SmtpFrom
				}

				b, err := json.Marshal(gCfg.AlertmanagerConfig)
				require.NoError(t, err)

				require.Equal(t, string(b), cfg.RawConfig)
			}

			// Receiver names should be unique.
			var finalCfg definition.PostableApiAlertingConfig
			require.NoError(t, json.Unmarshal([]byte(cfg.RawConfig), &finalCfg))
			receiverNames := map[string]struct{}{}
			for _, rcv := range finalCfg.Receivers {
				_, ok := receiverNames[rcv.Name]
				require.False(t, ok, fmt.Sprintf("duplicate receiver name %q found in final configuration", rcv.Name))
				receiverNames[rcv.Name] = struct{}{}
			}

			// Ensure that configuration is deterministic. For example, ordering of receivers.
			// This is important for change detection.
			cfg2, err := createUsableGrafanaConfig(am.logger, test.grafanaConfig, test.mimirConfig)
			require.NoError(t, err)

			require.Equal(t, cfg.RawConfig, cfg2.RawConfig)
		})
	}
}
