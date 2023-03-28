---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/aiven.svg"
brand_color: "#FF5733"
display_name: "Aiven"
short_name: "aiven"
description: "Steampipe plugin for Aiven."
og_description: "Query Aiven with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/aiven-social-graphic.png"
---

# Aiven + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Aiven](https://aiven.io) is the company that offers the best data infrastructure platform to the world.

For example:

```sql
select
  id,
  name,
  billing_enabled,
  tenant_id
from
  aiven_account;
```

```
+--------------+-------------+-----------------+-----------+
| id           | name        | billing_enabled | tenant_id |
+--------------+-------------+-----------------+-----------+
| a41fdc9a0625 | Turbot-test | false           | aiven     |
| a41fd3d3b210 | turbot      | false           | aiven     |
+--------------+-------------+-----------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/aiven/tables)**

## Get started

### Install

Download and install the latest Aiven plugin:

```bash
steampipe plugin install aiven
```

### Configuration

Installing the latest aiven plugin will create a config file (`~/.steampipe/config/aiven.spc`) with a single connection named `aiven`:

```hcl
connection "aiven" {
  plugin = "aiven"

  # You can connect to Aiven using one of options below:

  ## Use API Token authentication
  # `api_token` (required) - Create an authentication token in the Aiven Console for use with the Aiven CLI or API.
  # To learn more about using authentication tokens, refer to https://docs.aiven.io/docs/platform/concepts/authentication-tokens
  # Can also be set with the AIVEN_API_TOKEN environment variable.
  # api_token = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"

  # `user_agent` (optional) - Aiven user agent
  # Can also be set with the AIVEN_USER_AGENT environment variable.
  # user_agent = "aiven-go-client/v1.2.0"

  ## Use User authentication
  # `email` (required) - Aiven user email
  # Can also be set with the AIVEN_EMAIL environment variable.
  # email = "test@turbot.com"

  # `password` (required) - Aiven user password
  # Can also be set with the AIVEN_PASSWORD environment variable.
  # password = "test@123"

  # `user_agent` (optional) - Aiven user agent
  # Can also be set with the AIVEN_USER_AGENT environment variable.
  # user_agent = "aiven-go-client/v1.2.0"


  ## Use MFA User authentication
  # `email` (required) - Aiven user email
  # Can also be set with the AIVEN_EMAIL environment variable.
  # email = "test@turbot.com"

  # `password` (required) - Aiven user password
  # Can also be set with the AIVEN_PASSWORD environment variable.
  # password = "test@123"

  # `otp` (required) - Google authenticator OTP for Aiven
  # Can also be set with the AIVEN_OTP environment variable.
  # otp = "123456"

  # `user_agent` (optional) - Aiven user agent
  # Can also be set with the AIVEN_USER_AGENT environment variable.
  # user_agent = "aiven-go-client/v1.2.0"

  ## If no credentials are specified, the plugin will use Aiven CLI authentication
}
```

### Authentication token Credentials

You may specify the api token and user agent to authenticate:

- `api_token`(required): Specify the authentication token.
- `user_agent`(optional): Specify the user agent.

```hcl
connection "aiven_via_mfa_user" {
  plugin     = "aiven"
  email      = "test@turbot.com"
  password   = "turbot@123"
  user_agent = "aiven-go-client/v1.2.0"
}
```

### User Credentials

You may specify the email ID, password and user agent to authenticate:

- `email`(required): Specify the aiven email.
- `password`(required): Specify the aiven password.
- `user_agent`(optional): Specify the user agent.

```hcl
connection "aiven_via_mfa_user" {
  plugin     = "aiven"
  email      = "test@turbot.com"
  password   = "turbot@123"
  user_agent = "aiven-go-client/v1.2.0"
}
```

### MFA User Credentials

You may specify the email ID, password, OTP and user agent to authenticate:

- `email`(required): Specify the aiven email.
- `password`(required): Specify the aiven password.
- `otp`(required): Specify the google authenticator otp for aiven.
- `user_agent`(optional): Specify the user agent.

```hcl
connection "aiven_via_mfa_user" {
  plugin     = "aiven"
  email      = "test@turbot.com"
  password   = "turbot@123"
  otp        = "123456"
  user_agent = "aiven-go-client/v1.2.0"
}
```

### Credentials from Environment Variables

The Aiven plugin will use the Aiven environment variables to obtain credentials **only if other arguments (`api_token`, `user_agent`, `email`, `password`, etc..) are not specified** in the connection:

```sh
export AIVEN_API_TOKEN="oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4Cny"
export AIVEN_USER_AGENT="aiven-go-client/v1.2.0"
export AIVEN_EMAIL="test@turbot.com"
export AIVEN_PASSWORD="turbot@123"
export AIVEN_OTP="123456"
```

```hcl
connection "aiven" {
  plugin = "aiven"
}
```

### Aiven CLI

If no credentials are specified and the environment variables are not set, the plugin will use the active credentials from the Aiven CLI. You can run `avn user login` to set up these credentials.

```hcl
connection "aiven" {
  plugin = "aiven"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-aiven
- Community: [Slack Channel](https://steampipe.io/community/join)
