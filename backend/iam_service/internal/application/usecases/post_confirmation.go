package usecases

import (
	"context"

	"fivetrace.com/common/lib"
	"fivetrace.com/iam_service/internal/ports"
	"github.com/sirupsen/logrus"
)

type PostConfirmationUseCase struct {
	logger         *logrus.Entry
	paymentGateway ports.PaymentGatewayManager
}

func NewPostConfirmationUseCase(ctx context.Context, paymentGateway ports.PaymentGatewayManager) *PostConfirmationUseCase {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	return &PostConfirmationUseCase{
		logger:         logger,
		paymentGateway: paymentGateway,
	}
}

func (u *PostConfirmationUseCase) Execute(organizationName, email, sub string) error {
	logger := u.logger.WithField("email", email)

	logger.Info("Creating customer in payment gateway")

	_, err := u.paymentGateway.CreateCustomer(sub, email, organizationName)

	if err != nil {
		logger.WithError(err).Error("Failed to create customer in payment gateway")
		return err
	}

	return nil
}
