---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/aiven.svg"
brand_color: "#FF3554"
display_name: "Aiven"
short_name: "aiven"
description: "Steampipe plugin to query accounts, projects, teams, users and more from Aiven."
og_description: "Query Aiven with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/aiven-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Aiven + Steampipe

[Aiven](https://aiven.io) is the company that offers the best data infrastructure platform to the world.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List your Aiven accounts:

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

## Quick start

### Install

Download and install the latest Aiven plugin:

```sh
steampipe plugin install aiven
```

### Credentials

| Item        | Description                                                                                                                                                             |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Aiven requires an [API key](https://docs.aiven.io/docs/platform/howto/create_authentication_token) or `email` and `password` for all requests.                          |
| Permissions | API keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                               |
| Radius      | Each connection represents a single Aiven Installation.                                                                                                                 |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/aiven.spc`)<br />2. Credentials specified in environment variables, e.g., `AIVEN_TOKEN`. |

### Configuration

Installing the latest aiven plugin will create a config file (`~/.steampipe/config/aiven.spc`) with a single connection named `aiven`:

Configure your account details in `~/.steampipe/config/aiven.spc`:

```hcl
connection "aiven" {
  plugin = "aiven"

  # You can connect to Aiven using one of the options below:

  # Using API Key authentication
  # `api_key` (required) - Create an authentication token in the Aiven Console for use with the Aiven CLI or API.
  # To create an authentication token, refer to https://docs.aiven.io/docs/platform/howto/create_authentication_token
  # Can also be set with the AIVEN_TOKEN environment variable.
  # api_key = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"

  # Using User authentication (without 2FA)
  # email = "test@turbot.com"
  # password = "test@123"

  # If no credentials are specified, the plugin will use Aiven CLI authentication.
  # We recommend using API Key authentication for MFA user.
}
```

## Configuring Aiven Credentials

### Authentication Token Credentials

You may specify the API key to authenticate:

- `api_key`: Specify the authentication token.

```hcl
connection "aiven" {
  plugin   = "aiven"
  api_key  = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"
}
```

### User Credentials

You may specify the email ID and password to authenticate:

- `email`: Specify the aiven email.
- `password`: Specify the aiven password.

```hcl
connection "aiven" {
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

### Aiven CLI

If no credentials are specified and the environment variables are not set, the plugin will use the active credentials from the Aiven CLI. You can run `avn user login` to set up these credentials.

```hcl
connection "aiven" {
  plugin = "aiven"
}
```


