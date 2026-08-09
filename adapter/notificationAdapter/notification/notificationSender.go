package notification

type NotificationSender interface {
	Send(recipient string, message string)
}
