variable "aws_region" {
  description = "The AWS region to deploy resources in"
  type        = string
}

variable "stripe_secret_key" {
  description = "Stripe account secret key"
  type        = string
}