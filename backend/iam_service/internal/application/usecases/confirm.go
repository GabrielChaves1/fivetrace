package usecases

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"luminog.com/common/lib"
	"luminog.com/iam_service/internal/application/managers"
	"luminog.com/iam_service/internal/ports"
)

type ConfirmUseCase struct {
	ctx context.Context
	idp ports.IdentityProvider
	db  ports.AuthTokenTable
}

func NewConfirmUseCase(
	ctx context.Context,
	idp ports.IdentityProvider,
	db ports.AuthTokenTable,
) *ConfirmUseCase {
	return &ConfirmUseCase{
		ctx: ctx,
		idp: idp,
		db:  db,
	}
}

type ConfirmUseCaseError struct {
	Message    string
	StatusCode int
}

func (u *ConfirmUseCase) Execute(token string) *ConfirmUseCaseError {
	logger := lib.LoggerFromContext(u.ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	claims, err := managers.ParseJWT(token)

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	storedToken, err := u.db.GetToken(u.ctx, claims.Subject)

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

	logger.Info("jwt validated")

	err = u.idp.ConfirmSignUp(u.ctx, claims.Subject)

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    fmt.Sprintf("user confirmation error: %v", err),
			StatusCode: 400,
		}
	}

	err = u.db.DeleteToken(u.ctx, claims.Subject)

	if err != nil {
		return &ConfirmUseCaseError{
			Message:    fmt.Sprintf("error to delete token: %v", err),
			StatusCode: 400,
		}
	}

	return nil
}
