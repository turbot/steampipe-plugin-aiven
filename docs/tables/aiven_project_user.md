---
title: "Steampipe Table: aiven_project_user - Query Aiven Project Users using SQL"
description: "Allows users to query Aiven Project Users, providing detailed information about each user associated with a specific project."
---

# Table: aiven_project_user - Query Aiven Project Users using SQL

An Aiven Project User is a user that has been granted access to a specific project within the Aiven platform. This user has certain permissions and roles within the project, which can include managing services, viewing billing information, and modifying project settings. Each user is associated with an email address, which is used as their unique identifier within the project.

## Table Usage Guide

The `aiven_project_user` table provides insights into users within Aiven Projects. As a project manager or system administrator, explore user-specific details through this table, such as their email, access control rights, and associated metadata. Utilize it to uncover information about users, such as their permissions within a project, to ensure proper access control and security measures are in place.

## Examples

### Basic info
Explore which team members are associated with different projects in your organization. This can help in understanding the distribution of resources and roles within your projects.

```sql+postgres
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

```sql+sqlite
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
Explore which team members are part of the 'turbot' project. This can help in understanding the composition of the project team, including their roles and whether they are a billing contact.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have users with administrative access in various projects. This is beneficial for understanding project control distribution and identifying potential security risks.

```sql+postgres
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

```sql+sqlite
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
Gain insights into the accounts where no billing contact has been assigned to the user. This is beneficial for identifying potential gaps in account management and ensuring all users have appropriate billing contacts.

```sql+postgres
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

```sql+sqlite
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
  billing_contact = 0;
```

### List users who are not part of any team
Explore which users are not associated with any team, enabling you to identify potential areas of resource reallocation or team restructuring. This is particularly useful in assessing the overall distribution of users within your project.

```sql+postgres
select
  email,
  real_name,
  billing_contact,
  member_type,
  project_name,
  create_time
from
  aiven_project_user
where
  team_name = '';
```

```sql+sqlite
select
  email,
  real_name,
  billing_contact,
  member_type,
  project_name,
  create_time
from
  aiven_project_user
where
  team_name = '';
```