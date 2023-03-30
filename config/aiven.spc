connection "aiven" {
  plugin = "aiven"

  # You can connect to Aiven using one of options below:

  # Use API Key authentication
  # `api_key` (required) - Create an authentication token in the Aiven Console for use with the Aiven CLI or API.
  # To learn more about using authentication tokens, refer to https://docs.aiven.io/docs/platform/concepts/authentication-tokens
  # Can also be set with the AIVEN_TOKEN environment variable.
  # api_key = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"

  # Use User authentication
  # email = "test@turbot.com"
  # password = "test@123"

  # If no credentials are specified, the plugin will use Aiven CLI authentication.
  # We recommend using API Key authentication for MFA user.
}
