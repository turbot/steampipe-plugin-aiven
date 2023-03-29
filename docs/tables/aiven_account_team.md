# Table: aiven_account_team

Teams are specific to a single organization or organizational unit and cannot be shared between them. You can also use teams within organizations or organizational units to control access to projects for a group of users instead of specifying them per project. When you create a team, you choose which projects to associate it to and define the roles.

## Examples

### Basic info

```sql
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

```sql
select
  m ->> 'user_id' as user_id,
  m ->> 'real_name' as user_name,
  id as team_id,
  name as team_name
from
  aiven_account_team,
  jsonb_array_elements(members) as m
where
  name = 'dev';
```

### List teams which are not associated with any project

```sql
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

### List members who are admins in each account

```sql
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

### List members who have read_only access to projects

```sql
select
  m ->> 'user_id' as user_id,
  m ->> 'real_name' as user_name,
  id as team_id,
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
