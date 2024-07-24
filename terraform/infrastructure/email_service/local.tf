locals {
  domain_name                  = "luminog.com"
  sender_email                 = "noreply@${local.domain_name}"
  email_build_path = "../../backend/email_sender_service/build"
}