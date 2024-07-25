package ports

type PaymentGatewayCustomer struct {
	ID string
}


type PaymentGatewayManager interface {
	CreateCustomer(sub, email, organizationName string) (*PaymentGatewayCustomer, error)
}
