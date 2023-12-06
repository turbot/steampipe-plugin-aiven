---
title: "Steampipe Table: aiven_project_event_log - Query Aiven Project Event Logs using SQL"
description: "Allows users to query Aiven Project Event Logs, specifically the event history and related details of cloud projects, providing insights into project activities and potential issues."
---

# Table: aiven_project_event_log - Query Aiven Project Event Logs using SQL

Aiven Project Event Log is a resource within Aiven that allows you to monitor and track the history of events related to your cloud projects. It provides a detailed log of all project activities, including changes made, errors occurred, and actions taken. Aiven Project Event Log helps you stay informed about the status and performance of your projects and take appropriate actions when required.

## Table Usage Guide

The `aiven_project_event_log` table provides insights into the event logs within Aiven Project. As a DevOps engineer, explore event-specific details through this table, including event type, time, and associated metadata. Utilize it to uncover information about project activities, such as the changes made, errors occurred, and actions taken, to maintain the health and performance of your projects.

## Examples

### Basic info
Explore which events occurred in your project, who performed them, and when. This allows you to track changes and maintain a comprehensive log of project activities.

```sql+postgres
select
  actor,
  event_desc,
  event_type,
  project_name,
  time
from
  aiven_project_event_log;
```

```sql+sqlite
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
Discover who was responsible for removing a specific service, in this case 'redis-turbot', by exploring the event logs. This can be particularly useful for auditing purposes or for troubleshooting when a service unexpectedly disappears.

```sql+postgres
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

```sql+sqlite
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
Explore which users have established a new service integration within the past week. This is particularly useful to monitor recent system changes and maintain security compliance.

```sql+postgres
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

```sql+sqlite
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
  and time >= datetime('now', '-7 day') ;
```

### List users who powered down the services
Identify instances where users have powered down services, providing valuable insights into project management and service usage patterns.

```sql+postgres
select
  actor,
  event_desc,
  event_type,
  project_name,
  time
from
  aiven_project_event_log
where
  event_type = 'service_poweroff';
```

```sql+sqlite
select
  actor,
  event_desc,
  event_type,
  project_name,
  time
from
  aiven_project_event_log
where
  event_type = 'service_poweroff';
```