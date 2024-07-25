resource "aws_ssm_parameter" "stripe_secret_key" {
  name        = "/FiveTrace/StripeSecretKey"
  description = "Secret key da Stripe"
  type        = "SecureString"
  value       = var.stripe_secret_key

  tags = {
    environment = "dev"
  }
}