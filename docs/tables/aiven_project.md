# Table: aiven_project

Projects are collections of services and user permissions. Each project must have a unique name within an organization. You can group your services however you see fit.

## Examples

### Basic info

```sql
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

```sql
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

### List projects where services are not running

```sql
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

### List projects where default cloud is aws

```sql
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
