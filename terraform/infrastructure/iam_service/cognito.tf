resource "aws_cognito_user_pool" "cognito_user_pool" {
  name = "${local.project_name}-user-pool"

  username_attributes = ["email"]

  schema {
    name                = "email"
    attribute_data_type = "String"
    required            = true
    mutable             = true
  }

  schema {
    name                = "role"
    attribute_data_type = "String"
    mutable             = true
    required            = false
  }

  schema {
    name                = "name"
    attribute_data_type = "String"
    mutable             = true
    required            = false
  }

  schema {
    name                = "country"
    attribute_data_type = "String"
    mutable             = true
    required            = false
  }

  lifecycle {
    ignore_changes = [
      schema
    ]
  }
}

resource "aws_cognito_user_pool_domain" "pool_domain" {
  domain       = local.project_name
  user_pool_id = aws_cognito_user_pool.cognito_user_pool.id
}

resource "aws_cognito_resource_server" "cognito_resource_server" {
  user_pool_id = aws_cognito_user_pool.cognito_user_pool.id
  name         = "${local.project_name}-server"
  identifier   = "${local.project_name}-resource-server"

  scope {
    scope_name        = "read"
    scope_description = "for GET requests"
  }

  scope {
    scope_name        = "write"
    scope_description = "for POST requests"
  }
}

resource "aws_cognito_user_pool_client" "cognito_user_pool_client_dashboard" {
  name         = "${local.project_name}-dashboard"
  user_pool_id = aws_cognito_user_pool.cognito_user_pool.id

  generate_secret = true

  explicit_auth_flows = [
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_USER_SRP_AUTH"
  ]
}