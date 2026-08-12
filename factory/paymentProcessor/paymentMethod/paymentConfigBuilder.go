package paymentmethod

type PaymentConfigBuilder struct {
	creditCardNumber string
	payPalEmail      string
	cryptoAddress    string
}

func NewPaymentConfigBuilder() *PaymentConfigBuilder {
	return &PaymentConfigBuilder{}
}

func (this *PaymentConfigBuilder) CreditCardNumber(number string) *PaymentConfigBuilder {
	this.creditCardNumber = number
	return this
}

func (this *PaymentConfigBuilder) PayPalEmail(email string) *PaymentConfigBuilder {
	this.payPalEmail = email
	return this
}

func (this *PaymentConfigBuilder) CryptoAddress(address string) *PaymentConfigBuilder {
	this.cryptoAddress = address
	return this
}

func (this *PaymentConfigBuilder) Build() PaymentConfig {
	return PaymentConfig{
		creditCardNumber: this.creditCardNumber,
		payPalEmail:      this.payPalEmail,
		cryptoAddress:    this.cryptoAddress,
	}
}
