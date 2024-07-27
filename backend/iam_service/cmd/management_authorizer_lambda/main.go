package main

import (
	"context"
	"os"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/adapters"
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

func handler(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	contextFields := logrus.Fields{
		"requestID":         lambdaContext.AwsRequestID,
		"invokeFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	cognitoIDP := adapters.NewCognitoIdentityProvider(cognitoClient, &adapters.CognitoIdentityProviderConfig{
		ClientId:     os.Getenv("COGNITO_CLIENT_ID"),
		ClientSecret: os.Getenv("COGNITO_CLIENT_SECRET"),
		UserPoolId:   os.Getenv("COGNITO_USER_POOL_ID"),
	})

	effect := "Deny"

	managementAuthorizerUseCase := usecases.NewManagementAuthorizerUseCase(ctx, cognitoIDP)
	claims, isAllowed := managementAuthorizerUseCase.Execute(event.AuthorizationToken)

	if isAllowed {
		effect = "Allow"
	}

	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: claims.Subject,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action: []string{"execute-api:Invoke"},
					Effect: effect,
					Resource: []string{
						"*",
					},
				},
			},
		},
		Context: map[string]interface{}{
			"organization": claims.Organization,
		},
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
