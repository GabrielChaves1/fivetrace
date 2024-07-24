package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	
	"luminog.com/iam_serviceluminog.com/iam_service/internal/ports"
)

type SQSMessageQueue struct {
	client         *sqs.Client
	queueUrl       string
	messageGroupId string
}

func NewSQSMessageQueue(client *sqs.Client, queueUrl, messageGroupId string) ports.MessageQueue {
	return &SQSMessageQueue{client: client, queueUrl: queueUrl, messageGroupId: messageGroupId}
}

func (a *SQSMessageQueue) SendMessage(ctx context.Context, message []byte) error {
	stringifiedMsg := string(message)

	res, err := a.client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:            &stringifiedMsg,
		QueueUrl:               &a.queueUrl,
		MessageGroupId:         &a.messageGroupId,
		MessageDeduplicationId: aws.String(fmt.Sprintf("%d", time.Now().UnixNano())),
	})

	if err != nil {
		return err
	}

	fmt.Println(res.MessageId)

	return nil
}
