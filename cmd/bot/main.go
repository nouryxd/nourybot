package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/lyx0/nourybot/pkg/common"
	"github.com/sirupsen/logrus"
)

type Application struct {
	TwitchClient *twitch.Client
	Logger       *logrus.Logger
}

func main() {
	cfg := config.New()
	lgr := logrus.New()

	tc := twitch.NewClient(cfg.TwitchUsername, cfg.TwitchOauth)
	app := &Application{
		TwitchClient: tc,
		Logger:       lgr,
	}
	// Received a PrivateMessage (normal chat message), pass it to
	// the handler who checks for further action.
	app.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		app.Logger.Infof("%s", message)

		app.handlePrivateMessage(message)
	})

	// Received a WhisperMessage (Twitch DM), pass it to
	// the handler who checks for further action.
	// twitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
	// 	app.handleWhisperMessage(message)
	// })

	// Successfully connected to Twitch so we log a message with the
	// mode we are currently running in..
	app.TwitchClient.OnConnect(func() {
		app.Logger.Infof("Successfully connected to Twitch Servers in %s mode!", cfg.Environment)
		common.Send("nourylul", "xd", app.TwitchClient)
	})
	app.TwitchClient.Join("nourylul")

	err := app.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}
}
