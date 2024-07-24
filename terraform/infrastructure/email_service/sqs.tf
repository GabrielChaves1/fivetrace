resource "aws_sqs_queue" "email_sender_queue" {
  name       = "EmailSenderQueue.fifo"
  fifo_queue = true
  content_based_deduplication = true

  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.email_sender_dead_letter_queue.arn
    maxReceiveCount      = 5
  })
}

resource "aws_sqs_queue" "email_sender_dead_letter_queue" {
  name       = "EmailSenderDeadLetterQueue.fifo"
  fifo_queue = true
}

resource "aws_sqs_queue_redrive_allow_policy" "email_sender_dead_letter_queue_allow_policy" {
  queue_url = aws_sqs_queue.email_sender_dead_letter_queue.id

  redrive_allow_policy = jsonencode({
    redrivePermission = "byQueue",
    sourceQueueArns   = [aws_sqs_queue.email_sender_queue.arn]
  })
}