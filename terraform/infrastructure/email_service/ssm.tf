resource "aws_ssm_parameter" "sender_email_identity" {
  name  = "/Luminog/SenderEmailIdentity"
  type  = "SecureString"
  value = "Luminog <${aws_sesv2_email_identity.sender_email_identity.email_identity}>"
}