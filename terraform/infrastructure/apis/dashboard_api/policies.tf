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

data "aws_iam_policy_document" "signup_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "cognito-idp:AdminCreateUser",
      "cognito-idp:AdminUpdateUserAttributes",
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

data "aws_iam_policy_document" "confirm_email_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "cognito-idp:AdminConfirmSignUp",
      "cognito-idp:AdminUpdateUserAttributes",
      "dynamodb:GetItem",
      "dynamodb:DeleteItem",
      "ssm:GetParameter"
    ]

    resources = [
      aws_ssm_parameter.stripe_secret_key.arn,
      var.cognito_user_pool.arn,
      var.dynamodb_auth_tokens_arn,
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "confirm_email_policy" {
  name_prefix = "ConfirmEmailPolicy"
  policy      = data.aws_iam_policy_document.confirm_email_policy.json
}

data "aws_iam_policy_document" "signin_policy" {
  version = "2012-10-17"
  statement {
    actions = [
      "cognito-idp:InitiateAuth",
    ]

    resources = [
      var.cognito_user_pool.arn,
    ]

    effect = "Allow"
  }
}

resource "aws_iam_policy" "signin_policy" {
  name_prefix = "SignInUserPolicy"
  policy      = data.aws_iam_policy_document.signin_policy.json
}