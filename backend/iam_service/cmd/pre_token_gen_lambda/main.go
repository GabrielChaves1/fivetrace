package main

import (
	"context"

	"fivetrace.com/common/lib"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/sirupsen/logrus"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithFields(logrus.Fields{
		"type":         "lambda.handler",
		"record":       contextFields,
		"organization": event.Request.UserAttributes["custom:organization"],
	})

	logger.Info("setting custom claims")

	event.Response.ClaimsOverrideDetails.ClaimsToSuppress = []string{
		"custom:customer_id",
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
