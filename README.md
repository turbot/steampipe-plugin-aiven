![image](https://hub.steampipe.io/images/plugins/turbot/aiven-social-graphic.png)

# Aiven Plugin for Steampipe

Use SQL to query projects, services, integration endpoints and more from Aiven.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/aiven)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/aiven/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-aiven/issues)

## Quick start

Download and install the latest Aiven plugin:

```bash
steampipe plugin install aiven
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/aiven#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/aiven#configuration).

### Configuring Aiven Credentials

Configure your account details in `~/.steampipe/config/aiven.spc`:

You may specify the API key to authenticate:

- `api_key`: Specify the authentication token.

```hcl
connection "aiven" {
  plugin   = "aiven"
  api_key  = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"
}
```

or you may specify the email ID and password to authenticate:

- `email`: Specify the aiven email.
- `password`: Specify the aiven password.

```hcl
connection "aiven" {
  plugin   = "aiven"
  email    = "test@turbot.com"
  password = "turbot@123"
}
```

or through environment variables

```sh
export AIVEN_TOKEN="oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4Cny"
```

or through the active credentials from the Aiven CLI. You can run `avn user login` to set up these credentials.

```hcl
connection "aiven" {
  plugin = "aiven"
}
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-aiven/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-aiven/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Aiven Plugin](https://github.com/turbot/steampipe-plugin-aiven/labels/help%20wanted)
