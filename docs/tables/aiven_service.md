# Table: aiven_service

Provides different Aiven services information.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
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

### List services with target cloud provider aws

```sql
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
