---
title: "Steampipe Table: aiven_account_team - Query Aiven Account Teams using SQL"
description: "Allows users to query Aiven Account Teams, specifically the details regarding team members, their roles and associated projects."
---

# Table: aiven_account_team - Query Aiven Account Teams using SQL

An Aiven Account Team is a group of users within the Aiven platform that have access to certain projects. The team structure allows for easy management of permissions and roles within the platform. Each team can have multiple users and projects associated with it.

## Table Usage Guide

The `aiven_account_team` table provides insights into the team structure within Aiven. As a DevOps engineer or a team lead, you can explore team-specific details through this table, including member details, their roles, and associated projects. Use this table to manage and audit team permissions, roles and project associations effectively.

## Examples

### Basic info
Explore which team was created or updated at what time within your Aiven account. This can help you track changes and manage your teams effectively.

```sql+postgres
select
  id,
  name,
  account_id,
  create_time,
  update_time
from
  aiven_account_team;
```

```sql+sqlite
select
  id,
  name,
  account_id,
  create_time,
  update_time
from
  aiven_account_team;
```

### List members of the dev team
Identify the members of your development team, including their user IDs and real names, to gain insights into team composition. This can be particularly useful in scenarios where understanding team structure and membership is important for project management or resource allocation.

```sql+postgres
select
  m ->> 'user_id' as user_id,
  m ->> 'real_name' as user_name,
  aiven_account_team.id as team_id,
  name as team_name
from
  aiven_account_team,
  jsonb_array_elements(members) as m
where
  name = 'dev';
```

```sql+sqlite
select
  json_extract(m.value, '$.user_id') as user_id,
  json_extract(m.value, '$.real_name') as user_name,
  aiven_account_team.id as team_id,
  name as team_name
from
  aiven_account_team,
  json_each(members) as m
where
  name = 'dev';
```

### List teams which are not associated with any project
Determine the teams within an account that are not associated with any projects. This can be useful for identifying unused resources or potential misconfigurations.

```sql+postgres
select
  id as team_id,
  name as team_name,
  account_id,
  create_time
from
  aiven_account_team
where
  projects = '[]';
```

```sql+sqlite
select
  id as team_id,
  name as team_name,
  account_id,
  create_time
from
  aiven_account_team
where
  projects = '[]';
```

### List the admins of each account
Discover the segments that have administrative access in each account. This could be useful for auditing purposes, ensuring that only authorized individuals have administrative privileges.

```sql+postgres
select
  a.name as account_name,
  m ->> 'user_id' as user_id,
  m ->> 'real_name' as user_name,
  t.id as team_id,
  t.name as team_name
from
  aiven_account as a,
  aiven_account_team as t,
  jsonb_array_elements(members) as m
where
  t.account_id = a.id
  and t.id = a.owner_team_id;
```

```sql+sqlite
select
  a.name as account_name,
  json_extract(m.value, '$.user_id') as user_id,
  json_extract(m.value, '$.real_name') as user_name,
  t.id as team_id,
  t.name as team_name
from
  aiven_account as a,
  aiven_account_team as t,
  json_each(a.members) as m
where
  t.account_id = a.id
  and t.id = a.owner_team_id;
```

### List members who have read_only access to projects
This query is useful for identifying team members who have only read access to certain projects. This can be beneficial in managing user permissions and ensuring that sensitive projects are adequately protected.

```sql+postgres
select
  m ->> 'user_id' as user_id,
  m ->> 'real_name' as user_name,
  aiven_account_team.id as team_id,
  name as team_name,
  p ->> 'project_name' as project_name,
  p ->> 'team_type' as team_type
from
  aiven_account_team,
  jsonb_array_elements(members) as m,
  jsonb_array_elements(projects) as p
where
  p ->> 'team_type' = 'read_only';
```

```sql+sqlite
select
  json_extract(m.value, '$.user_id') as user_id,
  json_extract(m.value, '$.real_name') as user_name,
  aiven_account_team.id as team_id,
  name as team_name,
  json_extract(p.value, '$.project_name') as project_name,
  json_extract(p.value, '$.team_type') as team_type
from
  aiven_account_team,
  json_each(members) as m,
  json_each(projects) as p
where
  json_extract(p.value, '$.team_type') = 'read_only';
```