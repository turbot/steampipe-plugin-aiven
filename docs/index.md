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

  # Use API Key authentication
  # `api_key` (required) - Create an authentication token in the Aiven Console for use with the Aiven CLI or API.
  # To learn more about using authentication tokens, refer to https://docs.aiven.io/docs/platform/concepts/authentication-tokens
  # Can also be set with the AIVEN_TOKEN environment variable.
  # api_key = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"

  # Use User authentication (without 2FA)
  # email = "test@turbot.com"
  # password = "test@123"

  # If no credentials are specified, the plugin will use Aiven CLI authentication.
  # We recommend using API Key authentication for MFA user.
}
```

### Authentication Token Credentials

You may specify the API key to authenticate:

- `api_key`: Specify the authentication token.

```hcl
connection "aiven_via_api_key" {
  plugin   = "aiven"
  api_key  = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"
}
```

### User Credentials

You may specify the email ID and password to authenticate:

- `email`: Specify the aiven email.
- `password`: Specify the aiven password.

```hcl
connection "aiven_via_user" {
  plugin   = "aiven"
  email    = "test@turbot.com"
  password = "turbot@123"
}
```

### Credentials from Environment Variables

The Aiven plugin will use the Aiven environment variable to obtain credentials **only if other arguments (`api_key`, `email`, `password`) are not specified** in the connection:

```sh
export AIVEN_TOKEN="oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4Cny"
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
