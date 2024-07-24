data "aws_iam_policy_document" "assume_role_policy" {
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

resource "aws_iam_role" "signup_role" {
  name_prefix        = "SignupRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "signup_policy_attachment" {
  role       = aws_iam_role.signup_role.name
  policy_arn = aws_iam_policy.signup_policy.arn
}

resource "aws_iam_role_policy_attachment" "logs_policy_attachment" {
  role       = aws_iam_role.signup_role.name
  policy_arn = aws_iam_policy.logs_policy.arn
}