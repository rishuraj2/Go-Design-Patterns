package notification

import "notificationadapter/adaptee"

type Teams struct {
	teamsWebhook *adaptee.TeamsWebhook
	webhookUrl   string
}

func NewTeamsAdapter(teamsWebhook *adaptee.TeamsWebhook, webhookUrl string) Teams {
	return Teams{
		teamsWebhook: teamsWebhook,
		webhookUrl:   webhookUrl,
	}
}

func (t Teams) Send(recipient string, message string) {
	t.teamsWebhook.SendCard(recipient, message, t.webhookUrl)
}
