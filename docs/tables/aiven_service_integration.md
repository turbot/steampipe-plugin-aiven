# Table: aiven_service_integration

Service integrations provide additional functionality and features by connecting different Aiven services together. This includes metrics integration which enables Aiven users to send advanced telemetry data to an Aiven for PostgreSQL database for metrics and visualize it in Aiven for Grafana. In addition to the metrics integration, log integration is supported that allows sending logs from any Aiven service to Aiven for OpenSearch.

## Examples

### Basic info

```sql
select
  service_integration_id,
  source_project,
  source_service,
  active,
  enabled,
  integration_type
from
  aiven_service_integration;
```

### List source details for the service integrations

```sql
select
  service_integration_id,
  source_project,
  source_service,
  source_endpoint_id,
  source_endpoint_name,
  source_service_type
from
  aiven_service_integration
where
  ;
```

### List source details for the service integrations

```sql
select
  service_integration_id,
  destination_project,
  destination_service,
  destination_endpoint_id,
  destination_endpoint_name,
  destination_service_type
from
  aiven_service_integration
where
  ;
```

### List disabled service integrations

```sql
select
  service_integration_id,
  source_project,
  source_service,
  active,
  enabled,
  integration_type
from
  aiven_service_integration
where
  not enabled;
```

### Get integration status for each service integrations

```sql
select
  service_integration_id,
  active,
  enabled,
  jsonb_pretty(integration_status -> 'state') as integration_status_state,
  integration_status ->> 'status_user_desc'
from
  aiven_service_integration;
```
