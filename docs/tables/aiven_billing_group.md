# Table: aiven_billing_group

Billing groups let you to set up common billing profiles in an organization that can then be used for any projects within that organization. This includes projects in organizational units under that organization. So, instead of entering the payment information every time you create a project, you can use the information saved in the billing group.

## Examples

### Basic info

```sql
select
  id,
  billing_group_name,
  account_id,
  card_id,
  billing_currency,
  city
from
  aiven_billing_group;
```

### Get billing group information for each project

```sql
select
  g.id as billing_group_id,
  g.billing_group_name,
  p.name as project_name,
  state,
  card_id,
  company
from
  aiven_billing_group as g,
  aiven_project as p
where
  g.id = p.billing_group_id;
```

### Get billing group information for turbot company

```sql
select
  id,
  billing_group_name,
  account_id,
  card_id,
  billing_currency,
  city
from
  aiven_billing_group
where
  company = 'turbot';
```
