---
title: "Steampipe Table: aiven_service_integration_endpoint - Query Aiven Service Integration Endpoints using SQL"
description: "Allows users to query Aiven Service Integration Endpoints, specifically the details of each integration endpoint within a project."
---

# Table: aiven_service_integration_endpoint - Query Aiven Service Integration Endpoints using SQL

Aiven Service Integration Endpoint is a feature within Aiven that allows you to connect Aiven services with external systems. It provides a centralized way to manage and configure connections for various Aiven resources, including databases, caches, and more. Aiven Service Integration Endpoint helps you establish secure and reliable connections between Aiven services and your external systems.

## Table Usage Guide

The `aiven_service_integration_endpoint` table provides insights into service integration endpoints within Aiven. As a DevOps engineer, explore endpoint-specific details through this table, including endpoint configurations, related services, and associated metadata. Utilize it to uncover information about endpoints, such as their configuration details, the services they connect to, and the status of these connections.

## Examples

### Basic info
Discover the segments that are linked to different project names and endpoint types. This can be useful in managing and organizing your service integration endpoints effectively.

```sql+postgres
select
  endpoint_id,
  endpoint_name,
  project_name,
  endpoint_type
from
  aiven_service_integration_endpoint;
```

```sql+sqlite
select
  endpoint_id,
  endpoint_name,
  project_name,
  endpoint_type
from
  aiven_service_integration_endpoint;
```

### List external service integration endpoints
Explore which external services are integrated with your project. This helps you maintain an overview of your project's dependencies and interactions.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that consist of service integration endpoints which are not linked to any service. This is beneficial in identifying unused resources, thereby aiding in efficient resource management and cost optimization.

```sql+postgres
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

```sql+sqlite
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
      json_extract(i.value, '$.source_endpoint_id') as endpoint_id
    from
      aiven_service,
      json_each(integrations) as i
    union
    select
      json_extract(i.value, '$.dest_endpoint_id') as endpoint_id
    from
      aiven_service,
      json_each(integrations) as i
  );
```

### List service integration endpoint settings
Explore the configuration details of service integration endpoints to understand their types and user settings. This can be useful to assess the elements within your service integration and identify potential areas for optimization or troubleshooting.

```sql+postgres
select
  endpoint_id,
  endpoint_name,
  endpoint_type,
  jsonb_pretty(user_config) as user_config
from
  aiven_service_integration_endpoint;
```

```sql+sqlite
select
  endpoint_id,
  endpoint_name,
  endpoint_type,
  user_config
from
  aiven_service_integration_endpoint;
```