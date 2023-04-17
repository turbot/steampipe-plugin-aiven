# Table: aiven_account

Aiven account represents organizations and organizational units. When you sign up for Aiven, an organization is created for you. You can use that to create a hierarchical structure that fits your needs. Organizational units can be nested within an organization. This gives you greater flexibility to organize your setup to meet your specific use cases.

## Examples

### Basic info

```sql
select
  id,
  name,
  billing_enabled,
  create_time,
  tenant_id
from
  aiven_account;
```

### List accounts with billing enabled

```sql
select
  id,
  name,
  billing_enabled,
  create_time,
  tenant_id
from
  aiven_account
where
  billing_enabled;
```

### List projects of each account

```sql
select
  a.name as account_name,
  p.name as project_name,
  available_credits,
  default_cloud,
  estimated_balance,
  payment_method
from
  aiven_account as a,
  aiven_project as p
where
  a.id = p.account_id;
```
