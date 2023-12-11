---
title: "Steampipe Table: aiven_service - Query Aiven Cloud Services using SQL"
description: "Allows users to query Aiven Cloud Services, specifically the details of various services provided by Aiven, offering insights into service configurations, plan details, and status."
---

# Table: aiven_service - Query Aiven Cloud Services using SQL

Aiven Cloud Services is a comprehensive platform that offers managed open-source data technologies, including databases, message queues, and search engines. It allows users to deploy and manage these technologies on a variety of cloud platforms, providing flexibility and scalability. Aiven Cloud Services is designed to handle the complexity of managing and maintaining these technologies, enabling users to focus on their core business.

## Table Usage Guide

The `aiven_service` table provides insights into the services offered within Aiven Cloud Services. As a DevOps engineer or a cloud architect, explore service-specific details through this table, including service configurations, plan details, and current status. Utilize it to manage and monitor your Aiven services, ensuring optimal configuration and performance.

## Examples

### Basic info
Analyze the settings to understand the status and type of each project within your Aiven service. This can help you assess the overall health of your projects and plan accordingly.

```sql+postgres
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service;
```

```sql+sqlite
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service;
```

### List premium services
Explore which services are categorized as premium in your project. This can help prioritize resource allocation and understand the cost structure better.

```sql+postgres
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  plan like 'Premium%';
```

```sql+sqlite
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  plan like 'Premium%';
```

### List services which are not running
Identify instances where certain services are not currently operational. This is useful in pinpointing areas for troubleshooting or optimizing system performance.

```sql+postgres
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  state <> 'RUNNING';
```

```sql+sqlite
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  state <> 'RUNNING';
```

### List services with termination protection disabled
Determine the areas in which services have termination protection disabled. This is useful in identifying potential vulnerabilities and ensuring all services are adequately protected.

```sql+postgres
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  not termination_protection;
```

```sql+sqlite
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  termination_protection = 0;
```

### List services with target cloud provider aws
Explore services hosted on AWS by examining their names, project affiliations, states, plans, types, and creation times. This helps in managing and monitoring services specific to AWS, aiding in efficient resource allocation and project planning.

```sql+postgres
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  cloud_name like 'aws%';
```

```sql+sqlite
select
  name,
  project_name,
  state,
  plan,
  type,
  create_time
from
  aiven_service
where
  cloud_name like 'aws%';
```