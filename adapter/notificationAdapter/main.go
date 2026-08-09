package main

import (
	"notificationadapter/adaptee"
	"notificationadapter/notification"
)

func main() {
	slack := adaptee.NewSlackClient()
	teams := adaptee.NewTeamsWebhook()
	discord := adaptee.NewDiscordBot()

	slackAdapter := notification.NewSlackAdapter(&slack)
	teamsAdapter := notification.NewTeamsAdapter(&teams, "12345")
	discordAdapter := notification.NewDiscordAdapter(&discord, 123456)

	msg := "Hi Everyone"
	recipient := "ABC"

	slackAdapter.Send(recipient, msg)
	teamsAdapter.Send(recipient, msg)
	discordAdapter.Send(recipient, msg)
}
