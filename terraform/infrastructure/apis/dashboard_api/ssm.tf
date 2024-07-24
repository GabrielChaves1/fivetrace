resource "aws_ssm_parameter" "frontend_url" {
  name        = "front_url"
  description = "Front-end URL"
  type        = "String"
  value       = "http://localhost:5173"

  tags = {
    environment = "dev"
  }
}