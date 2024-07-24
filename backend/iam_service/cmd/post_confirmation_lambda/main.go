package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/sirupsen/logrus"
	"luminog.com/common/lib"
	"luminog.com/iam_service/internal/adapters"
	"luminog.com/iam_service/internal/application/usecases"

	"github.com/stripe/stripe-go/v79"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmationRequest) (events.CognitoEventUserPoolsPostConfirmation, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	email := event.UserAttributes["email"]

	stripePaymentGatewayManager := adapters.NewStripePaymentGatewayManager()

	postConfirmationUseCase := usecases.NewPostConfirmationUseCase(ctx, stripePaymentGatewayManager)
	postConfirmationUseCase.Execute(email)

	return events.CognitoEventUserPoolsPostConfirmation{}, nil
}

func main() {
	lambda.Start(handler)
}

func init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}
