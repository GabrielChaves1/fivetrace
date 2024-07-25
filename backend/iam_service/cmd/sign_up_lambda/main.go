package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"

	"luminog.com/common/lib"
	"luminog.com/iam_service/internal/adapters"
	"luminog.com/iam_service/internal/application/dtos"
	"luminog.com/iam_service/internal/application/usecases"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	cognitoClient *cognitoidentityprovider.Client
	dynamoClient  *dynamodb.Client
	sqsClient     *sqs.Client
	frontendURL   string
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	var request dtos.SignupDTO
	err := json.Unmarshal([]byte(event.Body), &request)
	if err != nil {
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

	authTokenTable := adapters.NewDynamoTokenTable(dynamoClient, "auth_tokens")
	sqsQueue := adapters.NewSQSMessageQueue(sqsClient, os.Getenv("SQS_QUEUE_URL"), "email")
	signupUseCase := usecases.NewSignupUseCase(ctx, cognitoIDP, authTokenTable, sqsQueue, frontendURL)

	fail := signupUseCase.Execute(request.Email, request.Password, request.OrganizationName, request.Country)

	if fail != nil {
		return events.APIGatewayProxyResponse{
			Body:       fail.Message,
			StatusCode: fail.StatusCode,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "User signed up",
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

	logger.Info("Initializing AWS SQS Client")
	sqsClient = sqs.NewFromConfig(cfg)

	logger.Info("Initializing AWS SSM")
	smmClient := ssm.NewFromConfig(cfg)

	res, err := smmClient.GetParameter(context.Background(), &ssm.GetParameterInput{
		Name: aws.String("front_url"),
	})

	if err != nil {
		panic("unable to get parameter, " + err.Error())
	}

	logger.Info("Setting environment variables", *res.Parameter.Value)

	frontendURL = *res.Parameter.Value
}

func main() {
	lambda.Start(handler)
}
