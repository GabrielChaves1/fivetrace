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

resource "aws_iam_policy" "logs_policy" {
  name_prefix = "CloudwatchLogsPolicy"
  policy      = data.aws_iam_policy_document.logs_policy.json
}