package bot

import (
	twitch "github.com/gempir/go-twitch-irc/v2"
	cfg "github.com/lyx0/nourybot/pkg/config"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	twitchClient *twitch.Client
	cfg          *cfg.Config
}

func NewBot(cfg *cfg.Config) *Bot {

	log.Info("fn Newbot")
	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)

	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
	}
}

func (b *Bot) Connect() error {
	log.Info("fn Connect")
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

func (b *Bot) Say(channel string, message string) {
	b.twitchClient.Say(channel, message)
}

func (b *Bot) OnPrivateMessage(callback func(message *twitch.PrivateMessage)) {
	log.Info(callback)
}

func (b *Bot) Join(channel string) {
	log.Info("fn Join")
	b.twitchClient.Join(channel)
}
