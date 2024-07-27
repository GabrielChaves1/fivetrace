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

resource "aws_iam_role_policy_attachment" "signup_logs_policy" {
  role       = aws_iam_role.signup_role.name
  policy_arn = aws_iam_policy.logs_policy.arn
}

resource "aws_iam_role" "confirm_email_role" {
  name_prefix = "ConfirmEmailRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "confirm_email_attachment" {
  role = aws_iam_role.confirm_email_role.name
  policy_arn = aws_iam_policy.confirm_email_policy.arn
}

resource "aws_iam_role_policy_attachment" "confirm_email_logs_policy" {
  role       = aws_iam_role.confirm_email_role.name
  policy_arn = aws_iam_policy.logs_policy.arn
}

resource "aws_iam_role" "signin_role" {
  name_prefix = "SignInUserRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "signin_attachment" {
  role = aws_iam_role.signin_role.name
  policy_arn = aws_iam_policy.signin_policy.arn
}

resource "aws_iam_role_policy_attachment" "signin_logs_policy" {
  role       = aws_iam_role.signin_role.name
  policy_arn = aws_iam_policy.logs_policy.arn
}