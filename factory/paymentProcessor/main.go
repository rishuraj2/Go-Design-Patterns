package main

import (
	"fmt"
	paymentmethod "paymentprocessor/paymentMethod"
)

func main() {
	config := paymentmethod.NewPaymentConfigBuilder().PayPalEmail("abc@some.mailcom").Build()
	method, err := paymentmethod.NewPaymentMethod("paypal", config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !method.ValidatePayment() {
		fmt.Println("validation failed!")
		return
	}

	method.ProcessPayment(10000.0)
	fmt.Println(method.GetReceipt())
}