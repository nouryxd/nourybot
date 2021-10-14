package bot

import (
	"time"

	twitch "github.com/gempir/go-twitch-irc/v2"
	cfg "github.com/lyx0/nourybot/pkg/config"
	"github.com/lyx0/nourybot/pkg/handlers"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	twitchClient *twitch.Client
	cfg          *cfg.Config
	Uptime       time.Time
}

func NewBot(cfg *cfg.Config) *Bot {

	log.Info("fn Newbot")
	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)

	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
		Uptime:       time.Now(),
	}
}

func (b *Bot) Connect() error {
	log.Info("fn Connect")
	cfg := cfg.LoadConfig()

	b.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		handlers.HandlePrivateMessage(message, b.twitchClient, cfg, b.Uptime)
	})

	b.twitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		handlers.HandleWhisperMessage(message, b.twitchClient)
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

func (b *Bot) Join(channel string) {
	log.Info("fn Join")
	b.twitchClient.Join(channel)
}
