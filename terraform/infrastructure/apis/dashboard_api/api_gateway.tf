resource "aws_api_gateway_rest_api" "dashboard_api" {
  name = "DashboardAPI"
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

resource "aws_api_gateway_method_response" "signup_method_response" {
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  resource_id = aws_api_gateway_resource.signup_resource.id
  http_method = aws_api_gateway_method.signup_method.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin"  = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Headers" = true
  }

  response_models = {
    "application/json" = "Empty"
  }
}

resource "aws_api_gateway_integration_response" "signup_integration_response" {
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  resource_id = aws_api_gateway_resource.signup_resource.id
  http_method = aws_api_gateway_method.signup_method.http_method
  status_code = aws_api_gateway_method_response.signup_method_response.status_code

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
    "method.response.header.Access-Control-Allow-Methods" = "'POST'"
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type'"
  }
}

resource "aws_api_gateway_deployment" "signup_deployment" {
  depends_on = [
    aws_api_gateway_integration.signup_integration,
  ]
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  stage_name  = "dev"
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



resource "aws_api_gateway_deployment" "confirm_deployment" {
  depends_on = [
    aws_api_gateway_integration.confirm_integration,
  ]
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  stage_name  = "dev"
}

output "api_urls" {
  value = {
    signup = "${aws_api_gateway_deployment.signup_deployment.invoke_url}/${aws_api_gateway_resource.signup_resource.path_part}"
    confirm = "${aws_api_gateway_deployment.confirm_deployment.invoke_url}/${aws_api_gateway_resource.confirm_resource.path_part}"
  }
}