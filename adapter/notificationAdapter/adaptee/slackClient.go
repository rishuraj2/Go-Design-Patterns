package adaptee

import "fmt"

type SlackClient struct{}

func NewSlackClient() SlackClient {
	return SlackClient{}
}

func (s SlackClient) PostMessage(channel, text string, asBot bool) {
	fmt.Printf("[Slack] Channel: %s, Text: %s, AsBot: %t\n", channel, text, asBot)
}
