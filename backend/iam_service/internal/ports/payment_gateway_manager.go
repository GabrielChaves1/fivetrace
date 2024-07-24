package ports

type PaymentGatewayManager interface {
	CreateCustomer(email string) (string, error)
}