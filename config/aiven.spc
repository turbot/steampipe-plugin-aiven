connection "aiven" {
  plugin = "aiven"

  # You can connect to Aiven using one of options below:

  ## Use API Key authentication
  # `api_key` (required) - Create an authentication token in the Aiven Console for use with the Aiven CLI or API.
  # To learn more about using authentication tokens, refer to https://docs.aiven.io/docs/platform/concepts/authentication-tokens
  # Can also be set with the AIVEN_TOKEN environment variable.
  # api_key = "oGAxUvrjAdL3QBhWnaJI67Pc9P0rPDzDfhykzVfBYPlmvVH8WdJMKaeVKzcrl4CnyXpjiaKJCCNT+OkbpxfWdDNqwZPngS"

  # `user_agent` (optional) - Aiven user agent
  # user_agent = "aiven-go-client/v1.2.0"

  ## Use User authentication
  # `email` (required) - Aiven user email
  # email = "test@turbot.com"

  # `password` (required) - Aiven user password
  # password = "test@123"

  # `user_agent` (optional) - Aiven user agent
  # user_agent = "aiven-go-client/v1.2.0"


  ## Use MFA User authentication
  # `email` (required) - Aiven user email
  # email = "test@turbot.com"

  # `password` (required) - Aiven user password
  # password = "test@123"

  # `otp` (required) - Google authenticator OTP for Aiven
  # otp = "123456"

  # `user_agent` (optional) - Aiven user agent
  # user_agent = "aiven-go-client/v1.2.0"

  ## If no credentials are specified, the plugin will use Aiven CLI authentication
}
