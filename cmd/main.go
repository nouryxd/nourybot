package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/config"
	"github.com/lyx0/nourybot/pkg/db"
	"github.com/lyx0/nourybot/pkg/handlers"
	"github.com/sirupsen/logrus"
)

var nb *bot.Bot

func main() {

	// No need to join every channel when on development mode.
	runMode := flag.String("mode", "production", "Mode in which to run. (dev/production")
	flag.Parse()

	conf := config.LoadConfig()

	nb = &bot.Bot{
		TwitchClient: twitch.NewClient(conf.Username, conf.Oauth),
		MongoClient:  db.Connect(conf),
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

	if *runMode == "production" {
		// Production, joining all regular channels
		logrus.Info("[PRODUCTION]: Joining every channel.")
		db.InitialJoin(nb)

	} else if *runMode == "dev" {
		// Development, only join my two channels
		logrus.Info("[DEV]: Joining nouryqt and nourybot.")

		nb.TwitchClient.Join("nouryqt", "nourybot")
		nb.Send("nourybot", "[DEV] Badabing Badaboom Pepepains")
	}

	nb.TwitchClient.Connect()
}
