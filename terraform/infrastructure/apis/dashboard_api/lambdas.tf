resource "aws_lambda_function" "signup_lambda" {
  function_name = "SignupUser"
  description   = "this function registers a user"
  role          = aws_iam_role.signup_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.iam_build_path}/sign_up_lambda.zip"
  source_code_hash = filebase64sha256("${local.iam_build_path}/sign_up_lambda.zip")

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

resource "aws_lambda_permission" "signup_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.signup_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.dashboard_api.execution_arn}/*/*"
}

resource "aws_lambda_function" "confirm_email_lambda" {
  function_name = "ConfirmUserEmail"
  description   = "this function confirm user"
  role          = aws_iam_role.confirm_email_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.iam_build_path}/confirm_email_lambda.zip"
  source_code_hash = filebase64sha256("${local.iam_build_path}/confirm_email_lambda.zip")

  logging_config {
    log_format = "JSON"
  }

  environment {
    variables = {
      COGNITO_CLIENT_ID = var.cognito_user_pool_app_client.client_id
      COGNITO_CLIENT_SECRET = var.cognito_user_pool_app_client.client_secret
      COGNITO_USER_POOL_ID = var.cognito_user_pool.id
      STRIPE_SECRET_KEY = aws_ssm_parameter.stripe_secret_key.name
    }
  }
}

resource "aws_lambda_permission" "confirm_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.confirm_email_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.dashboard_api.execution_arn}/*/*"
}

resource "aws_lambda_function" "signin_lambda" {
  function_name = "SignInUser"
  description   = "this function sign-in user"
  role          = aws_iam_role.signin_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.iam_build_path}/sign_in_lambda.zip"
  source_code_hash = filebase64sha256("${local.iam_build_path}/sign_in_lambda.zip")

  logging_config {
    log_format = "JSON"
  }

  environment {
    variables = {
      COGNITO_CLIENT_ID = var.cognito_user_pool_app_client.client_id
      COGNITO_CLIENT_SECRET = var.cognito_user_pool_app_client.client_secret
      COGNITO_USER_POOL_ID = var.cognito_user_pool.id
    }
  }
}

resource "aws_lambda_permission" "signin_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.signin_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.dashboard_api.execution_arn}/*/*"
}