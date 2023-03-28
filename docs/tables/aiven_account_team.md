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

### List teams with admin access

```sql
select
  id,
  name,
  account_id,
  create_time,
  update_time
from
  aiven_account_team
where
  name = 'Account Owners';
```
