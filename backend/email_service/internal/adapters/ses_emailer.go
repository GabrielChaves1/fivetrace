package adapters

import (
	"context"

	"fivetrace.com/common/lib"

	"fivetrace.com/email_service/internal/ports"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/sirupsen/logrus"
)

type SESEmailer struct {
	ctx    context.Context
	client *ses.Client
	source string
}

func NewSESEmailer(ctx context.Context, client *ses.Client, source string) ports.Emailer {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type":   "adapter",
		"port":   "emailer",
		"source": source,
	})
	ctx = lib.WithLogger(ctx, logger)

	logger.Info("creating SESEmailer")
	return &SESEmailer{ctx, client, source}
}

func (e *SESEmailer) SendHTMLEmail(to, subject, htmlBody string) error {
	logger := lib.LoggerFromContext(e.ctx).WithFields(logrus.Fields{
		"to":      to,
		"subject": subject,
	})

	logger.Info("creating send email request input")
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data: &htmlBody,
				},
			},
			Subject: &types.Content{
				Data: &subject,
			},
		},
		Source: &e.source,
	}

	logger.Info("sending email request input")
	_, err := e.client.SendEmail(context.Background(), input)
	return err
}
