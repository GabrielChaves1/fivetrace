package ports

type PaymentGatewayManager interface {
	CreateCustomer(organizationName, email string) (string, error)
}