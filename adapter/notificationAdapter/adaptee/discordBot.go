package adaptee

import "fmt"

type DiscordBot struct{}

func NewDiscordBot() DiscordBot {
	return DiscordBot{}
}

func (d DiscordBot) SendMessage(channelId int64, content string, tts bool) {
	fmt.Printf("[Discord] ChannelId: %d, Content: %s, tts: %t\n", channelId, content, tts)
}
