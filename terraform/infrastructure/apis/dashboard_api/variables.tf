variable "cognito_user_pool" {
  type = object({
    id  = string
    arn = string
  })
}

variable "cognito_user_pool_app_client" {
  type = object({
    client_id     = string
    client_secret = string
  })
}

variable "email_sender_queue" {
  type = object({
    arn = string
    url = string
  })
}

variable "dynamodb_auth_tokens_arn" {
  type = string
}