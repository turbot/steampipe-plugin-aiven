# Table: aiven_service_integration_endpoint

Aiven supports integrating with a number of external systems. In order to connect your Aiven service with an external system you will need to define the endpoint.

## Examples

### Basic info

```sql
select
  endpoint_id,
  endpoint_name,
  project_name,
  endpoint_type
from
  aiven_service_integration_endpoint;
```

### List external service integration endpoints

```sql
select
  endpoint_id,
  endpoint_name,
  project_name,
  endpoint_type
from
  aiven_service_integration_endpoint
where
  endpoint_type like 'external%';
```

### List service integration endpoints which are not associated to any service

```sql
select
  endpoint_id,
  endpoint_name,
  project_name,
  endpoint_type
from
  aiven_service_integration_endpoint
where
  endpoint_id not in
  (
    select
      i ->> 'source_endpoint_id' as endpoint_id
    from
      aiven_service,
      jsonb_array_elements(integrations) as i
    union
    select
      i ->> 'dest_endpoint_id' as endpoint_id
    from
      aiven_service,
      jsonb_array_elements(integrations) as i
  );
```

### List service integration endpoint settings

```sql
select
  endpoint_id,
  endpoint_name,
  endpoint_type,
  jsonb_pretty(user_config) as user_config
from
  aiven_service_integration_endpoint;
```
