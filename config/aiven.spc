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
