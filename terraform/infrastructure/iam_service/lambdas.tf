resource "aws_lambda_function" "pre_token_gen_lambda" {
  function_name = "PreTokenGen"
  description   = "this function executes before token generation"
  role          = aws_iam_role.pre_token_gen_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.iam_build_path}/pre_token_gen_lambda.zip"
  source_code_hash = filebase64sha256("${local.iam_build_path}/pre_token_gen_lambda.zip")

  logging_config {
    log_format = "JSON"
  }
}

resource "aws_lambda_permission" "allow_cognito_pre_token_gen" {
  statement_id  = "AllowCognitoInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.pre_token_gen_lambda.function_name
  principal     = "cognito-idp.amazonaws.com"
  source_arn    = aws_cognito_user_pool.cognito_user_pool.arn
}