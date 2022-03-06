package main

import (
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/bot"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/lyx0/nourybot/internal/db"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()
	logrus.Info(cfg.Username, cfg.Oauth)

	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)
	now := time.Now()
	mongoClient := db.Connect(cfg)

	nb := &bot.Bot{
		TwitchClient: twitchClient,
		MongoClient:  mongoClient,
		Uptime:       now,
	}

	nb.TwitchClient.Join("nourybot")

	err := nb.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}

}
