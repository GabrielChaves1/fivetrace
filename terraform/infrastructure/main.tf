module "iam_service" {
  source = "./iam_service"
  aws_region = var.aws_region
  stripe_secret_key = var.stripe_secret_key
}

module "email_service" {
  source = "./email_service"
  aws_region = var.aws_region
}

module "apis" {
  source = "./apis"
  cognito_user_pool = module.iam_service.cognito_user_pool
  cognito_user_pool_app_client = module.iam_service.cognito_user_pool_app_client
  dynamodb_auth_tokens_arn = module.iam_service.dynamodb_auth_tokens_arn
  email_sender_queue = module.email_service.email_sender_queue
}