package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "User successfully signed in",
	}, nil
}
