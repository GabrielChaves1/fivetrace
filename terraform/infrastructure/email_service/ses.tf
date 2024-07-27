resource "aws_sesv2_configuration_set" "sender_configuration_set" {
  configuration_set_name = "FiveTraceEmailSender"

  reputation_options {
    reputation_metrics_enabled = true
  }
}

resource "aws_sesv2_email_identity" "sender_domain_identity" {
  configuration_set_name = aws_sesv2_configuration_set.sender_configuration_set.configuration_set_name
  email_identity         = local.domain_name
}

resource "aws_sesv2_email_identity" "sender_email_identity" {
  configuration_set_name = aws_sesv2_configuration_set.sender_configuration_set.configuration_set_name
  email_identity         = local.sender_email
}