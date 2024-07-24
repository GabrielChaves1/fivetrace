module "iam_service" {
  source = "./iam_service"
  aws_region = var.aws_region
}