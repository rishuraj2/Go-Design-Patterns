package orderservice

import "fmt"

type ShippingService struct {
	counter int
}

func NewShippingService() ShippingService {
	return ShippingService{
		counter: 5000,
	}
}

func (this *ShippingService) CreateShipment(productId string, address string) string {
	this.counter++
	id := fmt.Sprintf("TRK-%d", this.counter)
	fmt.Printf("Shipment created for %s to %s. Tracking: %s\n", productId, address, id)
	return id
}

func (this *ShippingService) CancleShipment(trackingNumber string) {
	fmt.Printf("Rolling back: Cancelling shipment %s\n", trackingNumber)
}
