package orderservice

import "fmt"

type PaymentService struct {
	counter int
}

func NewPaymentService() PaymentService {
	return PaymentService{
		counter: 4000,
	}
}

func (this *PaymentService) Charge(paymentMethod string, amount float64) string {
	if paymentMethod == "expired-card" {
		fmt.Println("failure")
		return ""
	}

	this.counter++
	id := fmt.Sprintf("TXN-%d", this.counter)

	fmt.Printf("Payment: Charged ₹%f to %s. Transaction: %s\n", amount, paymentMethod, id)
	return id
}

func (this *PaymentService) Refund(transactionId string) {
	this.counter--
	fmt.Printf("Rolling back: Refunding transaction %s\n", transactionId)
}
