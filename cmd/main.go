package main

import (
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/config"
	"github.com/lyx0/nourybot/pkg/db"
	"github.com/lyx0/nourybot/pkg/handlers"
)

var nb *bot.Bot

func main() {

	conf := config.LoadConfig()

	// db.InsertInitialData()

	nb = &bot.Bot{
		TwitchClient: twitch.NewClient(conf.Username, conf.Oauth),
		Uptime:       time.Now(),
	}

	nb.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {

		// If channelID is missing something must have gone wrong.
		channelID := message.Tags["room-id"]
		if channelID == "" {
			fmt.Printf("Missing room-id tag in message")
			return
		}

		// Don't act on bots own messages.
		if message.Tags["user-id"] == conf.BotUserId {
			return
		}

		// Forward the message for further processing.
		handlers.PrivateMessage(message, nb)
	})

	// nb.TwitchClient.Join("nouryqt", "nourybot")
	db.InitialJoin(nb)
	// nb.Send("nourybot", "HeyGuys")
	nb.TwitchClient.Connect()
}
