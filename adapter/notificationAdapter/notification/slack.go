package notification

import "notificationadapter/adaptee"

type SlackAdapter struct {
	slackClient *adaptee.SlackClient
}

func NewSlackAdapter(slackClient *adaptee.SlackClient) SlackAdapter {
	return SlackAdapter{
		slackClient: slackClient,
	}
}

func (s SlackAdapter) Send(recipient string, message string) {
	s.slackClient.PostMessage(recipient, message, true)
}
