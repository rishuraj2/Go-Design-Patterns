package paymentmethod

import (
	"fmt"
	"regexp"
)

type PayPal struct {
	email string
}

func init() {
	GetPaymentMethodStore().Register("paypal", PayPal{})
}

func (this PayPal) Build(config PaymentConfig) PaymentMethod {
	return PayPal{
		email: config.GetPayPalEmail(),
	}
}

func (this PayPal) ProcessPayment(amount float64) {
	fee := amount * 0.015
	fmt.Printf("Processing PayPal payment: ₹%.2f (fee: ₹%.2f)\n", amount, fee)
}

func (this PayPal) ValidatePayment() bool {
	fmt.Println("Validating PayPal account...")
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	return pattern.MatchString(this.email)
}

func (this PayPal) GetReceipt() string {
	return "PayPal: " + this.email
}
