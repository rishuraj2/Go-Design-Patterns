package orderfacade

import (
	"fmt"
	orderservice "onlinefoodorder/orderService"
)

type OrderFacade struct {
	inventory    *orderservice.InventoryService
	payment      *orderservice.PaymentService
	shipping     *orderservice.ShippingService
	notification *orderservice.NotificationService
	orderCounter int
}

func NewOrderService(inventory *orderservice.InventoryService, payment *orderservice.PaymentService, shipping *orderservice.ShippingService, notification *orderservice.NotificationService) OrderFacade {
	return OrderFacade{
		inventory:    inventory,
		payment:      payment,
		shipping:     shipping,
		notification: notification,
		orderCounter: 6000,
	}
}

func (this *OrderFacade) PlaceOrder(productId string, quantity int, paymentMethod string, amount float64, address string, email string) bool {
	fmt.Println("--- Processing Order ---")
	if !this.inventory.CheckStock(productId, quantity) {
		fmt.Println("Failure")
		return false
	}

	this.inventory.ReserveStock(productId, quantity)
	pid := this.payment.Charge(paymentMethod, amount)
	if pid == "" {
		this.inventory.ReleaseStock(productId, quantity)
		return false
	}

	sid := this.shipping.CreateShipment(productId, address)
	oid := fmt.Sprintf("ORD-%d", this.orderCounter)
	this.notification.SendOrderConfirmation(email, oid, sid)

	fmt.Println("--- Order Placed Successfully! ---")
	return true
}
