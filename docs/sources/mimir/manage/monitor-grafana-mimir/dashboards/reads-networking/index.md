---
aliases:
  - ../../../operators-guide/monitor-grafana-mimir/dashboards/reads-networking/
  - ../../../operators-guide/monitoring-grafana-mimir/dashboards/reads-networking/
  - ../../../operators-guide/visualizing-metrics/dashboards/reads-networking/
description: View an example Reads networking dashboard.
menuTitle: Reads networking
title: Grafana Mimir Reads networking dashboard
weight: 100
---

<!-- Note: This topic is mounted in the GEM documentation. Ensure that all updates are also applicable to GEM. -->

# Grafana Mimir Reads networking dashboard

The Reads networking dashboard shows receive and transmit bandwidth, in-flight requests, and TCP connections.
The dashboard isolates each service on the read path into its own section and displays the order in which a read request flows.

This dashboard requires [additional resources metrics](../../requirements/#additional-resources-metrics).

Use this dashboard for the following use cases:

- Monitor in-flight requests and TCP connections to identify any potential bottlenecks or connection issues between services.
- Visualize bandwidth to detect if any component is experiencing network saturation or unusual traffic patterns.

## Example

The following example shows a Reads networking dashboard from a demo cluster.

![Grafana Mimir reads networking dashboard](mimir-reads-networking.png)
