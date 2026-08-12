package paymentmethod

type PaymentConfig struct {
	creditCardNumber string
	payPalEmail      string
	cryptoAddress    string
}

func (this PaymentConfig) GetCreditCardNumber() string {
	return this.creditCardNumber
}

func (this PaymentConfig) GetPayPalEmail() string {
	return this.payPalEmail
}

func (this PaymentConfig) GetCryptoAddress() string {
	return this.cryptoAddress
}
