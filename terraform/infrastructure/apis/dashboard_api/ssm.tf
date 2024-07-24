resource "aws_ssm_parameter" "frontend_url" {
  name        = "front_url"
  description = "Front-end URL"
  type        = "SecureString"
  value       = "http://localhost:3000/"

  tags = {
    environment = "dev"
  }
}