package paymentmethod

import (
	"fmt"
	"regexp"
)

type CryptoPayment struct {
	walletID string
}

func init() {
	GetPaymentMethodStore().Register("crypto", CryptoPayment{})
}

func (this CryptoPayment) Build(config PaymentConfig) PaymentMethod {
	return CryptoPayment{
		walletID: config.GetCryptoAddress(),
	}
}

func (this CryptoPayment) ValidatePayment() bool {
	fmt.Println("Validating crypto wallet...")
	pattern := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)

	return pattern.MatchString(this.walletID)
}

func (this CryptoPayment) ProcessPayment(amount float64) {
	fee := amount * 0.005
	fmt.Printf("Processing crypto payment: ₹%.2f (fee: ₹%.2f)\n", amount, fee)
}

func (this CryptoPayment) GetReceipt() string {
	if len(this.walletID) < 10 {
		return "Wallet: Invalid"
	}

	return fmt.Sprintf(
		"Wallet: %s...%s",
		this.walletID[:6],
		this.walletID[len(this.walletID)-4:],
	)
}