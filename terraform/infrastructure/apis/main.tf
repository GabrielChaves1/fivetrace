module "dashboard_api" {
  source = "./dashboard_api"
  cognito_user_pool = var.cognito_user_pool
  cognito_user_pool_app_client = var.cognito_user_pool_app_client
  dynamodb_auth_tokens_arn = var.dynamodb_auth_tokens_arn
  email_sender_queue = var.email_sender_queue
  stripe_secret_key = var.stripe_secret_key
}