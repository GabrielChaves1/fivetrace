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
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/sirupsen/logrus"

	"luminog.com/common/lib"
	"luminog.com/email_service/internal/adapters"
	"luminog.com/email_service/internal/application/managers"
	"luminog.com/email_service/internal/application/usecases"
)

var (
	sesClient *ses.Client
	ssmClient *ssm.Client
)

func getEmailSenderIdentityParameter(logger *logrus.Entry) (emailSenderIdentityParameter string, err error) {
	emailSenderIdentityParameter = os.Getenv("SENDER_IDENTITY_PARAMETER")
	if emailSenderIdentityParameter == "" {
		panic("SENDER_IDENTITY_PARAMETER env var is not set!")
	}

	logger.WithField("emailSenderIdentityParameter", emailSenderIdentityParameter).Info("Searching email sender identity parameter in SSM")
	param, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(emailSenderIdentityParameter),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", errors.Join(errors.New("couldn't find email sender identity parameter in SSM"), err)
	}

	return *param.Parameter.Value, nil
}

func handler(ctx context.Context, event events.SQSEvent) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)

	contextFields := logrus.Fields{
		"requestId":          lambdaContext.AwsRequestID,
		"invokedFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	emailSenderIdentity, err := getEmailSenderIdentityParameter(logger)
	if err != nil {
		panic(err)
	}
	logger.Info("sender email identity ", emailSenderIdentity)

	sesEmailer := adapters.NewSESEmailer(ctx, sesClient, emailSenderIdentity)

	emailTemplateManager, err := managers.NewEmailTemplateManager(ctx)
	if err != nil {
		panic(err)
	}

	sendHTMLEmailUseCase, err := usecases.NewSendEmailUseCase(ctx, sesEmailer, emailTemplateManager)
	if err != nil {
		panic(err)
	}

	sqsConsumer := adapters.NewSQSQueueMessageConsumer(ctx, sendHTMLEmailUseCase)

	err = sqsConsumer.Consume(event)
	if err != nil {
		panic(err)
	}
}

func init() {
	logger := logrus.New().WithField("type", "lambda.init")
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	logger.Info("Initializing SES client")
	sesClient = ses.NewFromConfig(cfg)

	logger.Info("Initializing SSM client")
	ssmClient = ssm.NewFromConfig(cfg)
}

func main() {
	lambda.Start(handler)
}
