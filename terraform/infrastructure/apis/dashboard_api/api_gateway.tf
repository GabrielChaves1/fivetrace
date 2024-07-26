resource "aws_api_gateway_rest_api" "dashboard_api" {
  name        = "DashboardAPI"
  description = "API Gateway to website"
}

resource "aws_api_gateway_resource" "signup_resource" {
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  parent_id   = aws_api_gateway_rest_api.dashboard_api.root_resource_id
  path_part   = "signup"
}

resource "aws_api_gateway_method" "signup_method" {
  rest_api_id   = aws_api_gateway_rest_api.dashboard_api.id
  resource_id   = aws_api_gateway_resource.signup_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "signup_integration" {
  rest_api_id             = aws_api_gateway_rest_api.dashboard_api.id
  resource_id             = aws_api_gateway_resource.signup_resource.id
  http_method             = aws_api_gateway_method.signup_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.signup_lambda.invoke_arn
}

resource "aws_api_gateway_resource" "confirm_resource" {
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  parent_id   = aws_api_gateway_rest_api.dashboard_api.root_resource_id
  path_part   = "confirm"
}

resource "aws_api_gateway_method" "confirm_method" {
  rest_api_id   = aws_api_gateway_rest_api.dashboard_api.id
  resource_id   = aws_api_gateway_resource.confirm_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "confirm_integration" {
  rest_api_id             = aws_api_gateway_rest_api.dashboard_api.id
  resource_id             = aws_api_gateway_resource.confirm_resource.id
  http_method             = aws_api_gateway_method.confirm_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.confirm_email_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "dashboard_deployment" {
  depends_on = [
    aws_api_gateway_integration.signup_integration,
    aws_api_gateway_integration.confirm_integration,
  ]
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  stage_name  = "dev"
}

output "api_urls" {
  value = {
    signup  = "${aws_api_gateway_rest_api.dashboard_api.execution_arn}/dev/${aws_api_gateway_resource.signup_resource.path_part}"
    confirm = "${aws_api_gateway_rest_api.dashboard_api.execution_arn}/dev/${aws_api_gateway_resource.confirm_resource.path_part}"
  }
}