package bot

import (
	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/commands"
	cfg "github.com/lyx0/nourybot/pkg/config"
	"github.com/lyx0/nourybot/pkg/handlers"
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
	cfg := cfg.LoadConfig()
	b.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		handlers.HandleTwitchMessage(message, b.twitchClient, cfg)
	})

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

func (b *Bot) OnPrivateMessage(message *twitch.PrivateMessage) {
	log.Info("fn OnPrivateMessage")

	tc := b.twitchClient
	commands.HandleCommand(*message, tc)
}

func (b *Bot) Join(channel string) {
	log.Info("fn Join")
	b.twitchClient.Join(channel)
}
