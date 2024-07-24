resource "aws_lambda_function" "post_confirmation_lambda" {
  function_name = "PostConfirmation"
  role = aws_iam_role.post_confirmation_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.iam_build_path}/post_confirmation_lambda.zip"
  source_code_hash = filebase64sha256("${local.iam_build_path}/post_confirmation_lambda.zip")

  logging_config {
    log_format = "JSON"
  }

  environment {
    variables = {
      STRIPE_SECRET_KEY = aws_ssm_parameter.stripe_secret_key.name
    }
  }
}

resource "aws_lambda_permission" "allow_cognito" {
  statement_id  = "AllowCognitoInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.post_confirmation_lambda.function_name
  principal     = "cognito-idp.amazonaws.com"
  source_arn    = aws_cognito_user_pool.cognito_user_pool.arn
}