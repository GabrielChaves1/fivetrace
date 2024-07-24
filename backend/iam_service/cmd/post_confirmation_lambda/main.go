package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmationRequest) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// Do something with the event here

	

	return events.CognitoEventUserPoolsPostConfirmation{}, nil
}

func main() {
	lambda.Start(handler)
}