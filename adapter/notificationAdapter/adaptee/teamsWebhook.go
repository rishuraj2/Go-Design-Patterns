package adaptee

import "fmt"

type TeamsWebhook struct{}

func NewTeamsWebhook() TeamsWebhook {
	return TeamsWebhook{}
}

func (t TeamsWebhook) SendCard(title, body, webhookUrl string) {
	fmt.Printf("[Teams] Title: %s, Body: %s, WebhookUrl: %s\n", title, body, webhookUrl)
}
