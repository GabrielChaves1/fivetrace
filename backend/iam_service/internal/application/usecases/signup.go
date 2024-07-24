package usecases

import (
	"context"
	"encoding/json"

	"luminog.com/common/lib"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"

	"luminog.com/iam_service/internal/ports"
	"luminog.com/iam_service/internal/utils"
)

var logBaseFields = logrus.Fields{
	"type": "usecase",
}

type SignupUseCase struct {
	ctx              context.Context
	idp              ports.IdentityProvider
	tokenTable       ports.AuthTokenTable
	emailSenderQueue ports.MessageQueue
	frontendUrl      string
}

func NewSignupUseCase(
	ctx context.Context,
	idp ports.IdentityProvider,
	tokenTable ports.AuthTokenTable,
	emailSenderQueue ports.MessageQueue,
	frontendUrl string,
) *SignupUseCase {
	return &SignupUseCase{
		ctx:              ctx,
		idp:              idp,
		tokenTable:       tokenTable,
		emailSenderQueue: emailSenderQueue,
		frontendUrl:      frontendUrl,
	}
}

type SignupUseCaseError struct {
	Message    string
	StatusCode int
}

func (u *SignupUseCase) Execute(email, password string) *SignupUseCaseError {
	logger := lib.LoggerFromContext(u.ctx).WithFields(logBaseFields)

	logger.Info("Starting signup process")

	sub, err := u.idp.SignUpUser(u.ctx, email, password)

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	token, err := utils.GenerateJWT(utils.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: sub,
		},
	})

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	err = u.tokenTable.PutToken(u.ctx, sub, token)

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	type EmailMessage struct {
		EmailType string                 `json:"emailType"`
		To        string                 `json:"to"`
		Data      map[string]interface{} `json:"data"`
	}

	emailMessage := EmailMessage{
		EmailType: "signup-confirm",
		To:        email,
		Data: map[string]interface{}{
			"link": u.frontendUrl + "/confirm?token=" + token,
		},
	}

	userJSON, err := json.Marshal(emailMessage)

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	err = u.emailSenderQueue.SendMessage(u.ctx, userJSON)

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	logger.Info("Sending confirmation link to email")

	return nil
}
