resource "aws_lambda_function" "email_sender_lambda" {
  function_name = "EmailSender"
  description   = "This function receives messages from a Queue and send e-mails different types of e-mails"
  role          = aws_iam_role.email_sender_role.arn

  runtime = "provided.al2023"
  handler = "bootstrap"

  filename         = "${local.email_build_path}/sqs_consumer_lambda.zip"
  source_code_hash = filebase64sha256("${local.email_build_path}/sqs_consumer_lambda.zip")

  environment {
    variables = {
      SENDER_IDENTITY_PARAMETER = aws_ssm_parameter.sender_email_identity.name
    }
  }

  logging_config {
    log_format = "JSON"
  }
}

resource "aws_lambda_event_source_mapping" "email_sender_queue_event_mapping" {
  event_source_arn = aws_sqs_queue.email_sender_queue.arn
  function_name    = aws_lambda_function.email_sender_lambda.arn
}