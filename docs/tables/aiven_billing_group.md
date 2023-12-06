---
title: "Steampipe Table: aiven_billing_group - Query Aiven Billing Groups using SQL"
description: "Allows users to query Billing Groups on Aiven, specifically the details of each billing group, providing insights into their configurations and usage."
---

# Table: aiven_billing_group - Query Aiven Billing Groups using SQL

Aiven Billing Groups are a feature that allows users to manage their Aiven services' billing under a single entity. It enables users to group their services together for easier invoicing and cost allocation. Billing Groups can be utilized across multiple projects, providing flexibility in managing costs and resources.

## Table Usage Guide

The `aiven_billing_group` table provides insights into Billing Groups within Aiven. As a financial analyst or a DevOps engineer, explore details of each billing group through this table, including their configurations, associated services, and usage. Utilize it to uncover information about billing groups, such as their costs, associated projects, and the allocation of resources.

## Examples

### Basic info
Explore which billing groups are associated with specific accounts and cards. This can be useful for understanding the financial structure of your organization and managing costs effectively.

```sql+postgres
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

```sql+sqlite
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
Explore which billing groups are associated with each project to better manage and understand your organization's financial landscape. This is useful for financial planning, cost allocation, and understanding the overall billing structure of your projects.

```sql+postgres
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

```sql+sqlite
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

### Get billing group information for a specific company
Analyze the billing group details for a specific company to gain insights into the billing currency and location. This can be useful for financial planning and geographical analysis.

```sql+postgres
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
  company = 'Dunder Mifflin';
```

```sql+sqlite
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
  company = 'Dunder Mifflin';
```