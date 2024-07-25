package usecases

import (
	"context"
	"fmt"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/application/managers"
	"fivetrace.com/iam_service/internal/ports"
	"github.com/sirupsen/logrus"
)

type ConfirmUseCase struct {
	logger *logrus.Entry
	idp    ports.IdentityProvider
	db     ports.AuthTokenTable
}

func NewConfirmUseCase(
	ctx context.Context,
	idp ports.IdentityProvider,
	db ports.AuthTokenTable,
) *ConfirmUseCase {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	return &ConfirmUseCase{
		logger: logger,
		idp:    idp,
		db:     db,
	}
}

type ConfirmUseCaseError struct {
	Message    string
	StatusCode int
}

func (u *ConfirmUseCase) Execute(token string) *ConfirmUseCaseError {
	claims, err := managers.ParseJWT(token)

	u.logger.Info("Unmarshalling token")

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	u.logger.Info("Getting token from dynamodb")

	storedToken, err := u.db.GetToken(context.Background(), claims.Subject)

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	if token != storedToken {
		return &ConfirmUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	u.logger.Info("Confirming user")

	err = u.idp.ConfirmSignUp(context.Background(), claims.Subject)

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    fmt.Sprintf("user confirmation error: %v", err),
			StatusCode: 400,
		}
	}

	u.logger.Info("Deleting token from dynamodb")

	err = u.db.DeleteToken(context.Background(), claims.Subject)

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    fmt.Sprintf("error to delete token: %v", err),
			StatusCode: 400,
		}
	}

	return nil
}
