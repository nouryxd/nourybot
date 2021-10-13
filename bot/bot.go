package bot

import (
	twitch "github.com/gempir/go-twitch-irc/v2"
	cfg "github.com/lyx0/nourybot/config"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	twitchClient *twitch.Client
	cfg          *cfg.Config
	log          *log.Logger
}

func (b *Bot) newTwitchClient() *twitch.Client {
	twitchClient := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return twitchClient
}

func NewBot(cfg *cfg.Config, log *log.Logger, twitchClient *twitch.Client) *Bot {
	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
		log:          log,
	}
}
