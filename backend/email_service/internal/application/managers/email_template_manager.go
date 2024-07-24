package managers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"

	"github.com/sirupsen/logrus"
	"luminog.com/common/lib"
	"luminog.com/email_service/internal/domain"
)

var logBaseFields = logrus.Fields{
	"type": "manager",
}

type EmailTemplateManager struct {
	ctx                 context.Context
	confirmLinkTemplate *template.Template
}

func NewEmailTemplateManager(ctx context.Context) (templateManager *EmailTemplateManager, err error) {
	logger := lib.LoggerFromContext(ctx).WithFields(logBaseFields)

	logger.Info("creating EmailTemplateManager")

	logger.Info("loading ConfirmLink template")
	confirmLinkTemplate, err := loadConfirmLinkTemplate()
	if err != nil {
		err = errors.Join(errors.New("couldn't create ConfirmLink template"), err)
		logger.Error(err)
		return nil, err
	}

	return &EmailTemplateManager{ctx, confirmLinkTemplate}, nil
}

func (a *EmailTemplateManager) formatConfirmLinkEmail(data map[string]interface{}) (subject string, body string, err error) {
	logger := lib.LoggerFromContext(a.ctx).WithFields(logBaseFields).WithFields(logrus.Fields{
		"data": data,
	})

	subject = "Confirmação de E-mail"

	logger.Info("validating ConfirmLink data fields")
	if _, ok := data["link"]; !ok {
		err := fmt.Errorf("invalid data format for ConfirmLink template, must have a 'link' property")
		logger.Error(err)
		return "", "", err
	}

	logger.Info("parsing ConfirmLink data")
	var bodyBuffer bytes.Buffer
	if err := a.confirmLinkTemplate.Execute(&bodyBuffer, data); err != nil {
		err := errors.Join(fmt.Errorf("couldn't parse email data"), err)
		logger.Error(err)
		return "", "", err
	}

	body = bodyBuffer.String()

	return subject, body, nil
}

func (a *EmailTemplateManager) FormatEmail(emailTemplateData domain.EmailTemplateData) (subject string, body string, err error) {
	logger := lib.LoggerFromContext(a.ctx).WithFields(logBaseFields).WithFields(logrus.Fields{
		"type": emailTemplateData.Type.String(),
		"data": emailTemplateData.Data,
	})

	logger.Info("formatting email")

	switch emailTemplateData.Type {
	case domain.ConfirmLink:
		return a.formatConfirmLinkEmail(emailTemplateData.Data)
	default:
		return "", "", fmt.Errorf("invalid e-mail template")
	}
}

func loadConfirmLinkTemplate() (templateInstance *template.Template, err error) {
	templateContent := `
    <html lang="pt-br">
    <head>
        <meta charset="UTF-8">
        <link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">
        <style>
            body {
                font-family: 'Roboto', sans-serif;
            }
        </style>
    </head>
    <body>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
            <tr>
                <td bgcolor="#ffffff" style="padding: 0px 0px 20px 0px;">
                    <div style="padding: 20px 0px 20px 0px;">
                        <p style="font-size: 16px; margin: 0;">Olá,</p>
                        <p style="font-size: 16px; margin: 0;">Estamos quase lá! Clique no link abaixo para confirmar seu cadastro:</p>
                    </div>
                    <p style="font-size: 16px; margin: 0;">
                        <a href="{{.link}}"
                            style="background-color: rgba(35,197,126,44%); border: solid 1px #26C57E; color: white; padding: 12px 20px; text-align: center; text-decoration: none; display: inline-block; border-radius: 4px;">Confirmar e-mail</a>
                    </p>
                </td>
            </tr>
            <tr>
                <td bgcolor="#f0f0f0" style="padding: 30px 30px 30px 30px;">
                    <p style="font-size: 14px; margin: 0;">Se você não solicitou acesso, ignore este email.
                    </p>
                </td>
            </tr>
        </table>
    </body>
    </html>
    `

	templateInstance, err = template.New("ConfirmLinkEmail").Parse(templateContent)
	if err != nil {
		return nil, err
	}

	return templateInstance, nil
}
