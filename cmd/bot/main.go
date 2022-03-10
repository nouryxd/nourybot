package main

import (
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/sirupsen/logrus"
)

type nourybot struct {
	twitchClient *twitch.Client
	uptime       time.Time
}

func main() {
	cfg := config.LoadConfig()
	logrus.Info(cfg.Username, cfg.Oauth)

	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)
	now := time.Now()

	nb := &nourybot{
		twitchClient: twitchClient,
		uptime:       now,
	}

	nb.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		nb.onPrivateMessage(message)
	})

	nb.twitchClient.OnConnect(nb.onConnect)

	err := nb.connect()
	if err != nil {
		panic(err)
	}

}
