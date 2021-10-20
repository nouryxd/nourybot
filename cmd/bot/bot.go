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

	// if text[0] == '.' || text[0] == '/' {
	// 	text = ". " + text
	// }

	if len(text) > 500 {
		firstMessage := text[0:499]
		secondMessage := text[499:]
		b.TwitchClient.Say(target, firstMessage)
		b.TwitchClient.Say(target, secondMessage)
		return
	}

	banned, reason := CheckMessage(text)
	if banned {
		b.TwitchClient.Say(target, reason)
	} else {
		b.TwitchClient.Say(target, text)
	}

}
