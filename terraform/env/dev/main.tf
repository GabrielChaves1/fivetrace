provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      "project"     = "Luminog"
      "environment" = "Development"
    }
  }
}

module "infrastructure" {
  source = "../../infrastructure"
  aws_region = var.aws_region
}