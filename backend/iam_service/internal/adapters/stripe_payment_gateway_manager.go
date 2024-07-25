package adapters

import (
	"fivetrace.com/iam_service/internal/ports"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
)

type StripePaymentGatewayManager struct{}

func NewStripePaymentGatewayManager() ports.PaymentGatewayManager {
	return &StripePaymentGatewayManager{}
}

func (a *StripePaymentGatewayManager) CreateCustomer(sub, email, organizationName string) (*ports.PaymentGatewayCustomer, error) {
	params := &stripe.CustomerParams{
		Name:  stripe.String(organizationName),
		Email: stripe.String(email),
		Metadata: map[string]string{
			"sub": sub,
		},
	}

	customer, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return &ports.PaymentGatewayCustomer{
		ID: customer.ID,
	}, nil
}
