# Table: aiven_project_event_log

Information about the events that have occurred in a project.

## Examples

### Basic info

```sql
select
  actor,
  event_desc,
  event_type,
  project_name,
  time
from
  aiven_project_event_log;
```

### Get user who deleted the service redis-turbot

```sql
select
  actor,
  event_desc,
  event_type,
  project_name,
  time
from
  aiven_project_event_log
where
  service_name = 'redis-turbot'
  and event_type = 'service_delete';
```

### List users who created service integration in last 7 days

```sql
select
  actor,
  event_desc,
  event_type,
  project_name,
  time
from
  aiven_project_event_log
where
  event_type = 'service_integration_create'
  and time >= (now() - interval '7' day) ;
```
