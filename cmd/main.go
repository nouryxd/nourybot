package main

import (
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/config"
	"github.com/lyx0/nourybot/pkg/handlers"
)

var nb *bot.Bot

func main() {
	conf := config.LoadConfig()

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

		// So that the bot doesn't repeat itself.
		if message.Tags["user-id"] == "596581605" {
			return
		}

		// Forward the message for further processing.
		handlers.PrivateMessage(message, nb)
	})

	nb.TwitchClient.Join("nouryqt", "nourybot")
	nb.Send("nourybot", "HeyGuys")
	nb.TwitchClient.Connect()
}
