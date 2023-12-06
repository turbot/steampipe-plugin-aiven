---
title: "Steampipe Table: aiven_account - Query Aiven Accounts using SQL"
description: "Allows users to query Aiven Accounts, providing detailed information about each account including its ID, name, owner and creation time."
---

# Table: aiven_account - Query Aiven Accounts using SQL

Aiven Account is a resource within Aiven's cloud-native database management system that represents the account of a user or organization. It provides a centralized way to manage and monitor the usage of various Aiven services, including databases, data pipelines, and more. Aiven Account helps users stay informed about the status and performance of their Aiven resources.

## Table Usage Guide

The `aiven_account` table provides insights into Aiven Accounts within Aiven's cloud-native database management system. As a database administrator or DevOps engineer, explore account-specific details through this table, including account ID, name, owner, and creation time. Utilize it to uncover information about accounts, such as their current status, usage of Aiven services, and the management of resources.

## Examples

### Basic info
Explore which Aiven accounts have billing enabled and when they were created to manage and optimize your financial resources effectively. This is useful for assessing cost efficiency and tracking the timeline of account creation.

```sql+postgres
select
  id,
  name,
  billing_enabled,
  create_time,
  tenant_id
from
  aiven_account;
```

```sql+sqlite
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
Explore which accounts have billing enabled. This can be useful for financial auditing or to identify accounts that are incurring charges.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  billing_enabled,
  create_time,
  tenant_id
from
  aiven_account
where
  billing_enabled = 1;
```

### List projects of each account
Uncover the details of each account's projects, including their available credits, default cloud, estimated balance, and payment method. This information can be useful for financial management and resource allocation.

```sql+postgres
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

```sql+sqlite
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