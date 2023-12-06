---
title: "Steampipe Table: aiven_project - Query Aiven Projects using SQL"
description: "Allows users to query Aiven Projects, providing access to details about each project such as its ID, name, and the time it was created."
---

# Table: aiven_project - Query Aiven Projects using SQL

Aiven Projects encapsulate a collection of services and are the main level of separation between different environments within Aiven. They provide a way to manage access controls and billing where the cost of all services under a project are combined into a single invoice. Aiven Projects are essential for managing and organizing resources in Aiven.

## Table Usage Guide

The `aiven_project` table provides insights into projects within Aiven. As a cloud engineer, explore project-specific details through this table, including the project's ID, name, and creation time. Utilize it to manage and organize your resources effectively within Aiven.

## Examples

### Basic info
Explore which Aiven projects have available credits and their estimated balances. This can be useful to manage budgets, by identifying which projects are nearing their credit limit and may require additional funding or cost control measures.

```sql+postgres
select
  name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project;
```

```sql+sqlite
select
  name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project;
```

### List projects with estimated bill greater than 50 USD
Explore which projects have an estimated bill greater than 50 USD to monitor and manage your budget effectively. This allows you to assess your financial commitments and plan for future expenses.

```sql+postgres
select
  name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project
where
  estimated_balance::float > 50;
```

```sql+sqlite
select
  name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project
where
  CAST(estimated_balance AS REAL) > 50;
```

### List projects where services are not running
Explore which projects have services that are not currently operational. This can help in identifying potential issues and ensuring all services are running as expected.

```sql+postgres
select
  p.name as project_name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project as p,
  aiven_service as s
where
  p.name = s.project_name
  and s.state <> 'RUNNING';
```

```sql+sqlite
select
  p.name as project_name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project as p,
  aiven_service as s
where
  p.name = s.project_name
  and s.state <> 'RUNNING';
```

### List projects where default cloud provider is aws
Discover the projects that have AWS as their default cloud provider. This can be useful to assess the distribution of your projects across different cloud platforms.

```sql+postgres
select
  name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project
where
  default_cloud like 'aws%';
```

```sql+sqlite
select
  name,
  account_id,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_project
where
  default_cloud like 'aws%';
```