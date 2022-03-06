package main

import (
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/bot"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/lyx0/nourybot/internal/db"
)

func main() {
	cfg := config.LoadConfig()

	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)
	now := time.Now()
	mongoClient := db.Connect(cfg)

	nb := &bot.Bot{
		TwitchClient: twitchClient,
		MongoClient:  mongoClient,
		Uptime:       now,
	}

	nb.TwitchClient.Join("nourybot")
	nb.TwitchClient.Connect()
}
