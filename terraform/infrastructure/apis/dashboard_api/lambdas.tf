resource "aws_lambda_function" "signup_lambda" {
  function_name = "SignupUser"
  description   = "this function registers a user"
  role          = aws_iam_role.signup_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.iam_build_path}/dashboard_sign_up_lambda.zip"
  source_code_hash = filebase64sha256("${local.iam_build_path}/dashboard_sign_up_lambda.zip")

  logging_config {
    log_format = "JSON"
  }

  environment {
    variables = {
      COGNITO_CLIENT_ID = var.cognito_user_pool_app_client.client_id
      COGNITO_CLIENT_SECRET = var.cognito_user_pool_app_client.client_secret
      COGNITO_USER_POOL_ID = var.cognito_user_pool.id
      SQS_QUEUE_URL = var.email_sender_queue.url
    }
  }
}