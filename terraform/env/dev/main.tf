provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      "project"     = "FiveTrace"
      "environment" = "Development"
    }
  }
}

module "infrastructure" {
  source = "../../infrastructure"
  aws_region = var.aws_region
  stripe_secret_key = var.stripe_secret_key
}