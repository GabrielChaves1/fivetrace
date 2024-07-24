output "cognito_user_pool" {
  description = "Cognito User Pool Details"
  value = {
    arn = aws_cognito_user_pool.cognito_user_pool.arn
    id = aws_cognito_user_pool.cognito_user_pool.id
  }
}

output "cognito_user_pool_app_client" {
  description = "Cognito Dashboard Application Client Details"
  value = {
    client_id = aws_cognito_user_pool_client.cognito_user_pool_client_dashboard.id
    client_secret = aws_cognito_user_pool_client.cognito_user_pool_client_dashboard.client_secret
  }
}

output "dynamodb_auth_tokens_arn" {
  description = "DynamoDB Auth Tokens Table ARN"
  value = aws_dynamodb_table.auth_tokens_table.arn
}