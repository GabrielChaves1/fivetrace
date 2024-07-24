provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      "project"     = "Luminog"
      "environment" = "Development"
    }
  }
}