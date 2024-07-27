package main

import (
	"context"
	"encoding/json"
	"os"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/adapters"
	"fivetrace.com/iam_service/internal/application/dtos"
	"fivetrace.com/iam_service/internal/application/usecases"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/sirupsen/logrus"
)

var (
	cognitoClient *cognitoidentityprovider.Client
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	var request dtos.SignInDTO
	err := json.Unmarshal([]byte(event.Body), &request)
	if err != nil {
		logger.WithError(err).Error("failed to unmarshal request body")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	cognitoIDP := adapters.NewCognitoIdentityProvider(cognitoClient, &adapters.CognitoIdentityProviderConfig{
		ClientId:     os.Getenv("COGNITO_CLIENT_ID"),
		ClientSecret: os.Getenv("COGNITO_CLIENT_SECRET"),
		UserPoolId:   os.Getenv("COGNITO_USER_POOL_ID"),
	})

	signInUseCase := usecases.NewSignInUseCase(ctx, cognitoIDP)
	tokens, err := signInUseCase.Execute(request)

	if err != nil {
		logger.WithError(err).Error("failed to execute sign in user")
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       tokens,
	}, nil
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

	logger.Info("Initializing AWS Cognito Client")
	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
}
