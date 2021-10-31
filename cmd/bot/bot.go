package bot

import (
	"time"

	twitch "github.com/gempir/go-twitch-irc/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bot struct {
	TwitchClient *twitch.Client
	MongoClient  *mongo.Client
	Uptime       time.Time
}

// Send checks the message against a banphrase api
// and also splits the message into two if the message
// is too long for a single twitch chat message.
func (b *Bot) Send(target, text string) {
	if len(text) == 0 {
		return
	}

	if text[0] == '.' || text[0] == '/' || text[0] == '!' {
		text = ":tf: " + text
	}

	// If a message is too long for a single twitch
	// message, split it into two messages.
	if len(text) > 500 {
		firstMessage := text[0:499]
		secondMessage := text[499:]

		b.TwitchClient.Say(target, firstMessage)
		b.TwitchClient.Say(target, secondMessage)

		return
	}

	// check the message for bad words before we say it
	messageBanned, banReason := CheckMessage(text)
	if messageBanned {
		// Bad message, replace message with a small
		// notice on why it's banned.
		b.TwitchClient.Say(target, banReason)
		return
	} else {
		// Message was okay.
		b.TwitchClient.Say(target, text)
		return
	}
}

// SkipChecking sends the message without banphrase check
func (b *Bot) SkipChecking(target, text string) {
	b.TwitchClient.Say(target, text)
}
