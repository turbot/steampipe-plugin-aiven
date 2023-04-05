# Table: aiven_account_authentication

You can authenticate directly with your email and password in the Aiven Console or use single sign-on through providers like GitHub, Google and Microsoft. You can also set up SAML authentication in Aiven with your company’s favorite authentication service.

## Examples

### Basic info

```sql
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

```sql
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

### List pending authentication methods

```sql
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

```sql
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

```sql
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
