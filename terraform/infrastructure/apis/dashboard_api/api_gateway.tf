resource "aws_api_gateway_rest_api" "dashboard_api" {
  name        = "DashboardAPI"
  description = "API Gateway to website"
}

resource "aws_api_gateway_deployment" "dashboard_deployment" {
  depends_on = [
    aws_api_gateway_integration.signup_integration,
    aws_api_gateway_integration.confirm_integration,
    aws_api_gateway_integration.signin_integration,
  ]
  rest_api_id = aws_api_gateway_rest_api.dashboard_api.id
  stage_name  = "dev"
}