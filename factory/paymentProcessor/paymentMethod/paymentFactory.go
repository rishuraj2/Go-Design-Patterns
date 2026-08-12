package paymentmethod


type PaymentMethod interface {
	ProcessPayment(amount float64)
	GetReceipt() string
	ValidatePayment() bool
	Build(config PaymentConfig) PaymentMethod
}

func NewPaymentMethod(methodName string, config PaymentConfig) (PaymentMethod, error) {
	method, err := GetPaymentMethodStore().Fetch(methodName)
	if err == nil {
		return method.Build(config), nil
	}

	return method, err
}
