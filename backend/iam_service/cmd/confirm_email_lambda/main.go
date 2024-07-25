package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/adapters"
	"fivetrace.com/iam_service/internal/application/usecases"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	cognitoClient *cognitoidentityprovider.Client
	dynamoClient  *dynamodb.Client
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	Token := event.QueryStringParameters["token"]

	cognitoIDP := adapters.NewCognitoIdentityProvider(cognitoClient, &adapters.CognitoIdentityProviderConfig{
		ClientId:     os.Getenv("COGNITO_CLIENT_ID"),
		ClientSecret: os.Getenv("COGNITO_CLIENT_SECRET"),
		UserPoolId:   os.Getenv("COGNITO_USER_POOL_ID"),
	})

	authTokenTable := adapters.NewDynamoTokenTable(dynamoClient, "auth_tokens")
	confirmUseCase := usecases.NewConfirmUseCase(ctx, cognitoIDP, authTokenTable)

	fail := confirmUseCase.Execute(Token)

	if fail != nil {
		return events.APIGatewayProxyResponse{
			Body:       fail.Message,
			StatusCode: fail.StatusCode,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "User confirmed successfully",
	}, nil
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

	logger.Info("Initializing AWS DynamoDB Client")
	dynamoClient = dynamodb.NewFromConfig(cfg)
}

func main() {
	lambda.Start(handler)
}
