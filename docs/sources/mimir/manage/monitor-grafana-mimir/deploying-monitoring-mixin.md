---
aliases:
  - ../../operators-guide/monitoring-grafana-mimir/deploying-monitoring-mixin/
  - ../../operators-guide/visualizing-metrics/deploying-monitor-mixin/
  - ../../operators-guide/monitor-grafana-mimir/deploying-monitoring-mixin/
description: Learn how to deploy the Grafana Mimir monitoring mixin.
menuTitle: Deploying the monitoring mixin
title: Deploying the Grafana Mimir monitoring mixin
weight: 20
---

<!-- Note: This topic is mounted in the GEM documentation. Ensure that all updates are also applicable to GEM. -->

# Deploying the Grafana Mimir monitoring mixin

Grafana Mimir exposes a `/metrics` endpoint returning Prometheus metrics. You can configure your Prometheus server to scrape Grafana Mimir or you can use the built-in functionality of the [Helm chart to automatically send these metrics to a remote](../collecting-metrics-and-logs/).
The endpoint is exposed on the Mimir HTTP server address / port which can be customized through `-server.http-listen-address` and `-server.http-listen-port` CLI flags or their respective YAML [config options](../../../configure/configuration-parameters/).

## Dashboards and alerts

Grafana Mimir is shipped with a comprehensive set of production-ready Grafana dashboards and alerts to monitor the state and health of a Mimir cluster.

Dashboards provide both a high-level and in-depth view of every aspect of a Grafana Mimir cluster.
You can take a look at all the available dashboards in [this overview](../dashboards/).

Alerts allow you to monitor the health of a Mimir cluster. For each alert, we provide detailed [runbooks](../../mimir-runbooks/) to further investigate and fix the issue.

The [requirements documentation](../requirements/) lists prerequisites for using the Grafana Mimir dashboards and alerts.

The [installation instructions](../installing-dashboards-and-alerts/) show available options to install Grafana Mimir dashboards and alerts.
