data "aws_iam_policy_document" "logs_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]

    effect = "Allow"
  }
}

data "aws_iam_policy_document" "signup_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "cognito-idp:AdminCreateUser",
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:DeleteItem",
      "sqs:SendMessage",
      "ssm:GetParameter"
    ]

    resources = [
      aws_ssm_parameter.frontend_url.arn,
      var.cognito_user_pool.arn,
      var.dynamodb_auth_tokens_arn,
      var.email_sender_queue.arn,
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "signup_policy" {
  name_prefix = "SignupPolicy"
  policy      = data.aws_iam_policy_document.signup_policy.json
}

resource "aws_iam_policy" "logs_policy" {
  name_prefix = "CloudwatchLogsPolicy"
  policy      = data.aws_iam_policy_document.logs_policy.json
}