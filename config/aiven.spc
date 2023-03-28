connection "aiven" {
  plugin = "aiven"

  # You can connect to Aiven using one of the options below:

  # API Token authentication
  # api_token(required) = "00000000-0000-0000-0000-000000000000"
  # user_agent(optional) = "aiven-go-client/v1.2.0"

  # User authentication
  # email(required) = "test@turbot.com"
  # password(required) = "test@123"
  # user_agent(optional) = "aiven-go-client/v1.2.0"


  # MFA User authentication
  # email(required) = "test@turbot.com"
  # password(required) = "test@123"
  # otp(required) = "123456"
  # user_agent(optional) = "aiven-go-client/v1.2.0"
}
