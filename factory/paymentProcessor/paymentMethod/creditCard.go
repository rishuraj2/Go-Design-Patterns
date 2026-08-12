package paymentmethod

import "fmt"

type CreditCard struct {
	number string
}

func init() {
	GetPaymentMethodStore().Register("credit card", CreditCard{})
}

func (this CreditCard) Build(config PaymentConfig) PaymentMethod {
	return CreditCard{
		number: config.GetCreditCardNumber(),
	}
}

func (this CreditCard) ProcessPayment(amount float64) {
	fee := amount * 0.025
	fmt.Printf("Processing credit card payment: ₹%.2f (fee: ₹%.2f)\n", amount, fee)
}

func (this CreditCard) ValidatePayment() bool {
	fmt.Println("Validating credit card...")
	return len(this.number) == 10
}

func (this CreditCard) GetReceipt() string {
	return fmt.Sprintf("Card ending in ****%s", this.number[5:])
}
