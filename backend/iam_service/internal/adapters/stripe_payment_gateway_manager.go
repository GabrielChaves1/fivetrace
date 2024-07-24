package adapters

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"luminog.com/iam_service/internal/ports"
)

type StripePaymentGatewayManager struct{}

func NewStripePaymentGatewayManager() ports.PaymentGatewayManager {
	return &StripePaymentGatewayManager{}
}

func (a *StripePaymentGatewayManager) CreateCustomer(email string) (string, error) {
	params := &stripe.CustomerParams{
		Name:  &email,
		Email: &email,
	}

	_, err := customer.New(params)

	if err != nil {
		return "", err
	}

	return "", nil
}
