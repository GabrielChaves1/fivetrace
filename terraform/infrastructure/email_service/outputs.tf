output "email_sender_queue" {
  value = {
    arn = aws_sqs_queue.email_sender_queue.arn
    url = aws_sqs_queue.email_sender_queue.url
  }
}