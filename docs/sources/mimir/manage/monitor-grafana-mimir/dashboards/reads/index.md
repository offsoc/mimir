---
aliases:
  - ../../../operators-guide/monitor-grafana-mimir/dashboards/reads/
  - ../../../operators-guide/monitoring-grafana-mimir/dashboards/reads/
  - ../../../operators-guide/visualizing-metrics/dashboards/reads/
description: View an example Reads dashboard.
menuTitle: Reads
title: Grafana Mimir Reads dashboard
weight: 90
---

<!-- Note: This topic is mounted in the GEM documentation. Ensure that all updates are also applicable to GEM. -->

# Grafana Mimir Reads dashboard

The Reads dashboard shows health metrics for the read path and object storage metrics for operations triggered by the read path.

The dashboard isolates each service on the read path into its own section and displays the order in which a read request flows.

Use this dashboard for the following use cases:

- Isolate components on the read path to help identify performance issues.
- Review health metrics for components such as the store-gateway, query-frontend, and querier.
- Asses whether object storage interactions are contributing to performance issues.

## Example

The following example shows a Reads dashboard from a demo cluster.

![Grafana Mimir reads dashboard](mimir-reads.png)
