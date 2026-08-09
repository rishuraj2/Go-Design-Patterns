package main

import (
	orderfacade "onlinefoodorder/orderFacade"
	orderservice "onlinefoodorder/orderService"
)

func main() {
	inventory := orderservice.NewInventoryService()
	payment := orderservice.NewPaymentService()
	shipping := orderservice.NewShippingService()
	notification := orderservice.NewNotificationService()

	facade := orderfacade.NewOrderService(&inventory, &payment, &shipping, &notification)

	_ = facade.PlaceOrder("SKU-001", 2, "credit-card", 49.99, "123 Main St", "user@example.com")

	_ = facade.PlaceOrder("SKU-001", 1, "expired-card", 29.99, "456 Oak Ave", "user@example.com")

}
