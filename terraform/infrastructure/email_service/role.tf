data "aws_iam_policy_document" "email_sender_assume_role_policy" {
  version = "2012-10-17"
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "email_sender_role" {
  name_prefix        = "EmailSender"
  assume_role_policy = data.aws_iam_policy_document.email_sender_assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "email_sender_role_send_email_policy_attachment" {
  role       = aws_iam_role.email_sender_role.name
  policy_arn = aws_iam_policy.send_email_policy.arn
}

resource "aws_iam_role_policy_attachment" "email_sender_role_send_logs_policy_attachment" {
  role       = aws_iam_role.email_sender_role.name
  policy_arn = aws_iam_policy.send_logs_policy.arn
}

resource "aws_iam_role_policy_attachment" "email_sender_role_queue_policy_attachment" {
  role       = aws_iam_role.email_sender_role.name
  policy_arn = aws_iam_policy.email_sender_queue_receive_message_policy.arn
}


resource "aws_iam_role_policy_attachment" "email_sender_role_parameter_read_policy_attachment" {
  role       = aws_iam_role.email_sender_role.name
  policy_arn = aws_iam_policy.email_sender_parameter_read_policy.arn
}