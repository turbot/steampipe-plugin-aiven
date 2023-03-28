![image](https://hub.steampipe.io/images/plugins/turbot/aiven-social-graphic.png)

# Aiven Plugin for Steampipe

Use SQL to query projects, services, integration endpoints and more from Aiven.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/aiven)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/aiven/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-aiven/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install aiven
```

Configure different methods for authentication in `~/.steampipe/config/aiven.spc`:

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

Run steampipe:

```shell
steampipe query
```

Query your account:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-aiven.git
cd steampipe-plugin-aiven
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/aiven.spc
```

Try it!

```
steampipe query
> .inspect aiven
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-aiven/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Aiven Plugin](https://github.com/turbot/steampipe-plugin-aiven/labels/help%20wanted)
