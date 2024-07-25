data "aws_iam_policy_document" "logs_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "ssm:GetParameter"
    ]

    resources = [
      "arn:aws:logs:*:*:*",
      aws_ssm_parameter.stripe_secret_key.arn
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "logs_policy" {
  name_prefix = "CloudwatchLogsPolicy"
  policy      = data.aws_iam_policy_document.logs_policy.json
}