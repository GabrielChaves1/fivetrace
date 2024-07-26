resource "aws_ssm_parameter" "sender_email_identity" {
  name  = "/FiveTrace/SenderEmailIdentity"
  type  = "SecureString"
  value = "FiveTrace <${aws_sesv2_email_identity.sender_email_identity.email_identity}>"
}