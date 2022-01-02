package main

import (
	"flag"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/go-delve/delve/pkg/config"
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
		// MongoClient:  db.Connect(conf),
		Uptime: time.Now(),
	}

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

}
