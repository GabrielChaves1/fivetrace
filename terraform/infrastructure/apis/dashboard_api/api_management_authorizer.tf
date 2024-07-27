resource "aws_api_gateway_authorizer" "management_authorizer" {
  rest_api_id    = aws_api_gateway_rest_api.dashboard_api.id
  name           = "ManagementAuthorizer"
  type           = "TOKEN"
  authorizer_uri = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${aws_lambda_function.management_authorizer_lambda.arn}/invocations"
  identity_source = "method.request.header.Authorization"
}