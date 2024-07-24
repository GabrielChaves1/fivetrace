package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"luminog.com/common/lib"
	"luminog.com/iam_service/internal/ports"
)

type PostConfirmationUseCase struct {
	ctx            context.Context
	paymentGateway ports.PaymentGatewayManager
}

func NewPostConfirmationUseCase(ctx context.Context, paymentGateway ports.PaymentGatewayManager) *PostConfirmationUseCase {
	return &PostConfirmationUseCase{
		ctx:            ctx,
		paymentGateway: paymentGateway,
	}
}

func (u *PostConfirmationUseCase) Execute(email string) error {
	logger := lib.LoggerFromContext(u.ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	logger.Info("creating customer in stripe")

	_, err := u.paymentGateway.CreateCustomer(email)

	if err != nil {
		return err
	}

	return nil
}
