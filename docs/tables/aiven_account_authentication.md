---
title: "Steampipe Table: aiven_account_authentication - Query Aiven Account Authentications using SQL"
description: "Allows users to query Aiven Account Authentications, providing insights into the authentication methods used across Aiven accounts."
---

# Table: aiven_account_authentication - Query Aiven Account Authentications using SQL

Aiven Account Authentication is a part of Aiven's account management system that handles the verification of identities and the management of access control. The authentication process ensures that the users are who they claim to be, thus providing a secure environment for data access and manipulation. It plays a crucial role in protecting sensitive data from unauthorized access and potential security threats.

## Table Usage Guide

The `aiven_account_authentication` table provides insights into the authentication methods used across Aiven accounts. As a security analyst, you can explore the details of each authentication method through this table, including the type of authentication, its status, and associated metadata. Utilize it to monitor and ensure the security of your Aiven accounts by verifying the authenticity of each user and their access controls.

## Examples

### Basic info
Explore which accounts are active and when they were created to manage access to your resources more effectively. This is useful for maintaining security and ensuring only authorized users have access.

```sql+postgres
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication;
```

```sql+sqlite
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication;
```

### List disabled authentication methods
Uncover the details of disabled authentication methods within your account. This can be useful for identifying potential security risks or areas for improvement within your authentication protocols.

```sql+postgres
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication
where
  not enabled;
```

```sql+sqlite
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication
where
  enabled = 0;
```

### List pending authentication methods
Discover the segments that have authentication methods still in the process of being configured. This is useful to ensure all methods are set up correctly and promptly for secure access.

```sql+postgres
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication
where
  state = 'pending_configuration';
```

```sql+sqlite
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication
where
  state = 'pending_configuration';
```

### List SAML authentication methods
Explore the SAML authentication methods in your Aiven account to determine which are enabled. This can be useful to identify potential security risks and maintain compliance with your organization's authentication policies.

```sql+postgres
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication
where
  type = 'saml';
```

```sql+sqlite
select
  id,
  name,
  enabled,
  state,
  type,
  create_time
from
  aiven_account_authentication
where
  type = 'saml';
```

### List expired SAML certificates
Assess the elements within your Aiven account authentication to identify expired SAML certificates. This is useful for maintaining security standards by ensuring all certificates are current and valid.

```sql+postgres
select
  id,
  name,
  enabled,
  state,
  type,
  create_time,
  saml_certificate_issuer,
  saml_certificate_not_valid_after
from
  aiven_account_authentication
where
  now() < saml_certificate_not_valid_after;
```

```sql+sqlite
select
  id,
  name,
  enabled,
  state,
  type,
  create_time,
  saml_certificate_issuer,
  saml_certificate_not_valid_after
from
  aiven_account_authentication
where
  datetime('now') < saml_certificate_not_valid_after;
```