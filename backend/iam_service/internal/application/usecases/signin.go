package usecases

import (
	"context"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/application/dtos"
	"fivetrace.com/iam_service/internal/application/validators"
	"fivetrace.com/iam_service/internal/ports"
	"github.com/sirupsen/logrus"
)

type SignInUseCase struct {
	logger *logrus.Entry
	idp    ports.IdentityProvider
}

func NewSignInUseCase(
	ctx context.Context,
	idp ports.IdentityProvider,
) *SignInUseCase {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	return &SignInUseCase{
		logger: logger,
		idp:    idp,
	}
}

func (u *SignInUseCase) Execute(signInDto dtos.SignInDTO) (string, error) {
	if err := validators.ValidateSignInDTO(signInDto); err != nil {
		u.logger.WithError(err).Error("invalid sign in data")
		return "", err
	}

	tokens, err := u.idp.SignInUser(context.Background(), signInDto.Email, signInDto.Password)

	if err != nil {
		u.logger.WithError(err).Error("failed to sign in user")
		return "", err
	}

	return tokens, nil
}
