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
	ctx               context.Context
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

	subject = "Seu acesso à experiência"

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
    <html lang="en">
			<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">
					<title>Email de Confirmação</title>
					<style>
							body {
									font-family: 'Roboto', sans-serif;
									background-color: #181818;
									margin: 0;
									padding: 0;
							}
							.container {
									max-width: 600px;
									margin: 0 auto;
									background-color: #161616;
									padding: 20px;
									border-radius: 4px;
									border: solid .1rem rgba(255,255,255,.1);
									border-top: none;
							}
							.content {
									text-align: center;
									padding: 20px 0;
									margin: 0 auto;
									max-width: 500px;
							}
							.content h1 {
									color: #E2E8F0;
							}
							.content p {
								color: rgba(226, 232, 240, .70);
							}
							.button-container {
									text-align: center;
									margin: 20px 0;
							}
							.button-container a {
									background-color: rgba(38, 197, 126, 0.45);
									border: solid .1rem rgb(38, 197, 126, .5);
									color: #ffffff;
									padding: 15px 25px;
									text-decoration: none;
									border-radius: 5px;
									font-size: 16px;
							}
							.footer {
									text-align: center;
									padding: 20px 0;
									color: rgba(226, 232, 240, 0.3);
									font-size: 12px;
							}
					</style>
			</head>
			<body>
					<div class="container">
							<div class="content">
									<h1>Confirme seu endereço de e-mail</h1>
									<p>Olá,</p>
									<p>Obrigado por se registrar em nossa plataforma. Para completar seu cadastro, por favor clique no botão abaixo para confirmar seu endereço de e-mail.</p>
							</div>
							<div class="button-container">
									<a href="link_de_confirmacao">Confirmar E-mail</a>
							</div>
							<div class="footer">
									<p>Se você não solicitou este e-mail, por favor ignore-o.</p>
									<p>&copy; 2024 Luminog. Todos os direitos reservados.</p>
							</div>
					</div>
			</body>
			</html>
    `

	templateInstance, err = template.New("ConfirmLinkEmail").Parse(templateContent)
	if err != nil {
		return nil, err
	}

	return templateInstance, nil
}
