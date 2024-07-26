package usecases

import (
	"context"
	"fmt"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/application/managers"
	"fivetrace.com/iam_service/internal/ports"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/sirupsen/logrus"
)

type ConfirmEmailUseCase struct {
	logger *logrus.Entry
	idp    ports.IdentityProvider
	db     ports.AuthTokenTable
	stripe ports.PaymentGatewayManager
}

func NewConfirmEmailUseCase(
	ctx context.Context,
	idp ports.IdentityProvider,
	db ports.AuthTokenTable,
	stripe ports.PaymentGatewayManager,
) *ConfirmEmailUseCase {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	return &ConfirmEmailUseCase{
		logger: logger,
		idp:    idp,
		db:     db,
		stripe: stripe,
	}
}

type ConfirmEmailUseCaseError struct {
	Message    string
	StatusCode int
}

func (u *ConfirmEmailUseCase) Execute(token string) *ConfirmEmailUseCaseError {
	claims, err := managers.ParseJWT(token)

	u.logger.Info("Unmarshalling token")

	if err != nil {
		return &ConfirmEmailUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	logger := u.logger.WithField("organization", claims.Organization)
	logger.Info("Getting token from dynamodb")

	storedToken, err := u.db.GetToken(context.Background(), claims.Subject)

	if err != nil {
		return &ConfirmEmailUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	if token != storedToken {
		return &ConfirmEmailUseCaseError{
			Message:    "invalid token",
			StatusCode: 400,
		}
	}

	logger.Info("confirming user in cognito")

	err = u.idp.ConfirmSignUp(context.Background(), claims.Subject)

	if err != nil {
		return &ConfirmEmailUseCaseError{
			Message:    fmt.Sprintf("user confirmation error: %v", err),
			StatusCode: 400,
		}
	}

	logger.Info("Deleting token from dynamodb")

	err = u.db.DeleteToken(context.Background(), claims.Subject)

	if err != nil {
		return &ConfirmEmailUseCaseError{
			Message:    fmt.Sprintf("error to delete token: %v", err),
			StatusCode: 400,
		}
	}

	logger.Info("creating customer in stripe")

	customer, err := u.stripe.CreateCustomer(claims.Subject, claims.Email, claims.Organization)

	if err != nil {
		return &ConfirmEmailUseCaseError{
			Message:    fmt.Sprintf("error to create customer in stripe: %v", err),
			StatusCode: 400,
		}
	}

	err = u.idp.UpdateUserAttributes(claims.Subject, []types.AttributeType{
		{
			Name:  aws.String("custom:customer_id"),
			Value: aws.String(customer.ID),
		},
	})

	if err != nil {
		return &ConfirmEmailUseCaseError{
			Message:    fmt.Sprintf("error to update attributes: %v", err),
			StatusCode: 400,
		}
	}

	return nil
}
