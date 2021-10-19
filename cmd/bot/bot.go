package bot

import (
	"time"

	twitch "github.com/gempir/go-twitch-irc/v2"
)

type Bot struct {
	TwitchClient *twitch.Client
	Uptime       time.Time
}

type Channel struct {
	Name string
}

func (b *Bot) Send(target, text string) {
	if len(text) == 0 {
		return
	}

	// if message[0] == '.' || message[0] == '/' {
	// 	message = ". " + message
	// }

	if len(text) > 500 {
		firstMessage := text[0:499]
		secondMessage := text[499:]
		b.TwitchClient.Say(target, firstMessage)
		b.TwitchClient.Say(target, secondMessage)
		return
	}

	b.TwitchClient.Say(target, text)
}
