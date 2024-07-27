package usecases

import (
	"context"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/application/managers"
	"fivetrace.com/iam_service/internal/domain"
	"fivetrace.com/iam_service/internal/ports"
	"github.com/sirupsen/logrus"
)

type ManagementAuthorizerUseCase struct {
	logger *logrus.Entry
	idp    ports.IdentityProvider
}

func NewManagementAuthorizerUseCase(
	ctx context.Context,
	idp ports.IdentityProvider,
) *ManagementAuthorizerUseCase {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	return &ManagementAuthorizerUseCase{
		logger: logger,
		idp:    idp,
	}
}

type ManagementAuthorizerUseCaseError struct {
	Message    string
	StatusCode int
}

func (u *ManagementAuthorizerUseCase) Execute(token string) (*managers.Claims, bool) {
	claims, err := managers.ParseJWT(token)

	if err != nil {
		u.logger.WithError(err).Error("failed to parse JWT")
		return nil, false
	}

	user, err := u.idp.GetUser(context.Background(), token)

	if err != nil {
		u.logger.WithError(err).Error("failed to get user")
		return nil, false
	}

	attributes := make(map[string]string)

	for _, attr := range user.UserAttributes {
		attributes[*attr.Name] = *attr.Value
	}

	if attributes["custom:role"] != domain.Manager.String() {
		u.logger.Error("user is not a manager")
		return nil, false
	}

	return claims, true
}
