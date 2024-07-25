package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/sirupsen/logrus"
	"luminog.com/common/lib"
	"luminog.com/iam_service/internal/adapters"
	"luminog.com/iam_service/internal/application/usecases"

	"github.com/stripe/stripe-go/v79"
)

var (
	ssmClient *ssm.Client
)

func getStripeSecretKey(logger *logrus.Entry) (stripeSecretKey string, err error) {
	secretKey := os.Getenv("STRIPE_SECRET_KEY")
	if secretKey == "" {
		panic("STRIPE_SECRET_KEY env var is not set!")
	}

	logger.Info("Searching stripe secret key in SSM")
	param, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(secretKey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", errors.Join(errors.New("couldn't find stripe secret key parameter in SSM"), err)
	}

	return *param.Parameter.Value, err
}

func handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	email := event.Request.UserAttributes["email"]
	organizationName := event.Request.UserAttributes["custom:organization"]

	stripePaymentGatewayManager := adapters.NewStripePaymentGatewayManager()

	postConfirmationUseCase := usecases.NewPostConfirmationUseCase(ctx, stripePaymentGatewayManager)
	err := postConfirmationUseCase.Execute(organizationName, email)

	if err != nil {
		logger.WithError(err).Error("Failed to execute post confirmation use case")
		return event, err
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}

func init() {
	logger := logrus.New().WithField("type", "lambda.init")
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	logger.Info("Initializing SSM client")
	ssmClient = ssm.NewFromConfig(cfg)

	stripeSecretKey, err := getStripeSecretKey(logger)

	if err != nil {
		panic(err)
	}

	stripe.Key = stripeSecretKey
}
