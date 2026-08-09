package notification

import "notificationadapter/adaptee"

type DiscordClient struct {
	discordClient *adaptee.DiscordBot
	channelId     int64
}

func NewDiscordAdapter(discord *adaptee.DiscordBot, channelId int64) DiscordClient {
	return DiscordClient{
		discordClient: discord,
	}
}

func (d DiscordClient) Send(recipient string, message string) {
	d.discordClient.SendMessage(d.channelId, message, false)
}
