package usecases

import (
	"context"
	"encoding/json"

	"luminog.com/common/lib"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"

	"luminog.com/iam_service/internal/application/dtos"
	"luminog.com/iam_service/internal/application/validators"
	"luminog.com/iam_service/internal/ports"
	"luminog.com/iam_service/internal/utils"
)

var logBaseFields = logrus.Fields{
	"type": "usecase",
}

type SignupUseCase struct {
	logger           *logrus.Entry
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
	logger := lib.LoggerFromContext(ctx).WithFields(logBaseFields)

	return &SignupUseCase{
		logger:           logger,
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

func (u *SignupUseCase) Execute(signupDto *dtos.SignupDTO) *SignupUseCaseError {
	if err := validators.ValidateSignupDTO(signupDto); err != nil {
		return &SignupUseCaseError{
			Message:    "invalid signup data",
			StatusCode: 400,
		}
	}

	logger := u.logger.WithField("organization_name", signupDto.OrganizationName).WithField("country", signupDto.Country)

	logger.Info("creating user in cognito")

	sub, err := u.idp.SignUpUser(signupDto.Email, signupDto.Password, signupDto.OrganizationName, signupDto.Country)

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	logger.Info("creating temp auth token")

	token, err := utils.GenerateJWT(utils.Claims{
		Email: signupDto.Email,
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

	logger.Info("saving temp auth token in dynamodb")

	err = u.tokenTable.PutToken(context.Background(), sub, token)

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
		EmailType: "confirm_link",
		To:        signupDto.Email,
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

	logger.Info("sending confirmation link to email")

	err = u.emailSenderQueue.SendMessage(context.Background(), userJSON)

	if err != nil {
		return &SignupUseCaseError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	return nil
}
