## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#39](https://github.com/turbot/steampipe-plugin-aiven/pull/39))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#39](https://github.com/turbot/steampipe-plugin-aiven/pull/39))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. 
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. 

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/install/steampipe.sh), as a [Postgres FDW](https://steampipe.io/install/postgres.sh), as a [SQLite extension](https://steampipe.io/install/sqlite.sh) and as a standalone [exporter](https://steampipe.io/install/export.sh). ([#15](https://github.com/turbot/steampipe-plugin-aiven/pull/15))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#15](https://github.com/turbot/steampipe-plugin-aiven/pull/15))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-aiven/blob/main/docs/LICENSE). ([#15](https://github.com/turbot/steampipe-plugin-aiven/pull/15))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server enacapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#14](https://github.com/turbot/steampipe-plugin-aiven/pull/14))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#8](https://github.com/turbot/steampipe-plugin-aiven/pull/8))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#6](https://github.com/turbot/steampipe-plugin-aiven/pull/6))
- Recompiled plugin with Go version `1.21`. ([#6](https://github.com/turbot/steampipe-plugin-aiven/pull/6))

## v0.0.1 [2023-04-17]

_What's new?_

- New tables added
  - [aiven_account](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_account)
  - [aiven_account_authentication](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_account_authentication)
  - [aiven_account_team](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_account_team)
  - [aiven_billing_group](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_billing_group)
  - [aiven_project](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_project)
  - [aiven_project_event_log](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_project_event_log)
  - [aiven_project_user](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_project_user)
  - [aiven_project_vpc](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_project_vpc)
  - [aiven_service](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_service)
  - [aiven_service_integration_endpoint](https://hub.steampipe.io/plugins/turbot/aiven/tables/aiven_service_integration_endpoint)
