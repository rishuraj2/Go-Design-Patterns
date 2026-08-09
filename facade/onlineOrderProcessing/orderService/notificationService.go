package orderservice

import "fmt"

type NotificationService struct{}

func NewNotificationService() NotificationService {
	return NotificationService{}
}

func (n *NotificationService) SendOrderConfirmation(email, orderId, trackingNumber string) {
	fmt.Printf("Notification: Order confirmation sent to %s. Order: %s, Tracking: %s\n", email, orderId, trackingNumber)
}
