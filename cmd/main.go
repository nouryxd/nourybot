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
	log "github.com/sirupsen/logrus"
)

var nb *bot.Bot

func main() {

	// runMode is either dev or production so that we don't join every channel
	// everytime we do a small test.
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

	// Depending on the mode we run in, join different channel.
	if *runMode == "production" {
		log.Info("[PRODUCTION]: Joining every channel.")

		// Production, joining all regular channels
		db.InitialJoin(nb)

	} else if *runMode == "dev" {
		log.Info("[DEV]: Joining nouryxd and nourybot.")

		// Development, only join my two channels
		nb.TwitchClient.Join("nouryxd", "nourybot")
		nb.Send("nourybot", "[DEV] Badabing Badaboom Pepepains")
	}

	nb.TwitchClient.Connect()
}
