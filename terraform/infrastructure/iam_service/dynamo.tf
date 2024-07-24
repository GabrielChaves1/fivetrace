resource "aws_dynamodb_table" "auth_tokens_table" {
  name           = "auth_tokens"
  hash_key       = "sub"
  write_capacity = 5
  read_capacity  = 5

  attribute {
    name = "sub"
    type = "S"
  }
}

resource "aws_dynamodb_table" "server_ownership_table" {
  name           = "server_ownership"
  hash_key       = "sub"
  range_key      = "server_id"
  write_capacity = 5
  read_capacity  = 5

  attribute {
    name = "sub"
    type = "S"
  }

  attribute {
    name = "server_id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "servers_data_table" {
  name           = "servers_data"
  hash_key       = "server_id"
  write_capacity = 5
  read_capacity  = 5

  attribute {
    name = "server_id"
    type = "S"
  }
}