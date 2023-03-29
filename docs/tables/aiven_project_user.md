# Table: aiven_project_user

Represents a user who has accepted membership in a project.

## Examples

### Basic info

```sql
select
  email,
  real_name,
  billing_contact,
  member_type,
  project_name,
  create_time,
  team_name
from
  aiven_project_user;
```

### List users of the turbot project

```sql
select
  email,
  real_name,
  billing_contact,
  member_type,
  project_name,
  create_time,
  team_name
from
  aiven_project_user
where
  project_name = 'turbot';
```

### List users with admin access in projects

```sql
select
  email,
  real_name,
  billing_contact,
  member_type,
  project_name,
  create_time,
  team_name
from
  aiven_project_user
where
  member_type = 'admin';
```

### List users with no billing contact

```sql
select
  email,
  real_name,
  billing_contact,
  member_type,
  project_name,
  create_time,
  team_name
from
  aiven_project_user
where
  not billing_contact;
```
