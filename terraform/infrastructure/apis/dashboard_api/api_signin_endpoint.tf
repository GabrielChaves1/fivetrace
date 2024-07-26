resource "aws_api_gateway_resource" "signin_resource" {
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  parent_id   = aws_api_gateway_rest_api.dashboard_api.root_resource_id
  path_part   = "login"
}

resource "aws_api_gateway_method" "signin_method" {
  rest_api_id   = aws_api_gateway_rest_api.dashboard_api.id
  resource_id   = aws_api_gateway_resource.signin_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "signin_integration" {
  rest_api_id             = aws_api_gateway_rest_api.dashboard_api.id
  resource_id             = aws_api_gateway_resource.signin_resource.id
  http_method             = aws_api_gateway_method.signin_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.signin_lambda.invoke_arn
}