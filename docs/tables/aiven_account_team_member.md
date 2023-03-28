# Table: aiven_account_team_member

Represents a team member, and it's associated team information.

**You must specify the account_id and team_id** in the where or join clause using the `account_id` and `team_id` column.

## Examples

### Basic info

```sql
select
  user_id,
  real_name,
  user_email,
  create_time,
  team_name
from
  aiven_account_team_member
where
  account_id = 'a41fdc9a0621'
  and team_id = 'at41fdc910261';
```

### List team members with admin access

```sql
select
  user_id,
  real_name,
  user_email,
  team_name,
  m.account_id as account_id
from
  aiven_account_team_member as m,
  aiven_account_team as t
where
  m.team_id = t.id
  and m.account_id = t.account_id
  and t.name = 'Account Owners';
```
