resource "aws_ssm_parameter" "frontend_url" {
  name        = "front_url"
  description = "Front-end URL"
  type        = "String"
  value       = "http://localhost:5173"

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "stripe_secret_key" {
  name        = "/FiveTrace/StripeSecretKey"
  description = "Secret key da Stripe"
  type        = "SecureString"
  value       = var.stripe_secret_key

  tags = {
    environment = "dev"
  }
}