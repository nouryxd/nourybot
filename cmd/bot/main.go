package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/sirupsen/logrus"
)

type Application struct {
	TwitchClient *twitch.Client
	Logger       *logrus.Logger
}

func main() {
	cfg := config.New()

	twitchClient := twitch.NewClient(cfg.TwitchUsername, cfg.TwitchOauth)

	app := Application{
		TwitchClient: twitchClient,
		Logger:       logrus.New(),
	}
	// Received a PrivateMessage (normal chat message), pass it to
	// the handler who checks for further action.
	app.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		app.handlePrivateMessage(message)
	})

	// Received a WhisperMessage (Twitch DM), pass it to
	// the handler who checks for further action.
	app.TwitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		app.handleWhisperMessage(message)
	})

	// Successfully connected to Twitch so we log a message with the
	// mode we are currently running in..
	app.TwitchClient.OnConnect(func() {
		app.Logger.Infof("Successfully connected to Twitch Servers in %s mode!", cfg.Environment)
		app.Send("nourylul", "xd")
	})

	app.TwitchClient.Join("nourylul")
	err := app.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}
}
