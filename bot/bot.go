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

func NewBot(cfg *cfg.Config, log *log.Logger, twitchClient *twitch.Client) *Bot {
	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
		log:          log,
	}
}

func (b *Bot) newTwitchClient() *twitch.Client {
	twitchClient := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return twitchClient
}

func (b *Bot) Connect() error {
	err := b.twitchClient.Connect()
	if err != nil {
		log.Error("Error Connecting from Twitch: ", err)
	}

	return err
}

func (b *Bot) Disconnect() error {
	err := b.twitchClient.Disconnect()
	if err != nil {
		log.Error("Error Disconnecting from Twitch: ", err)
	}

	return err
}
