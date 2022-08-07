package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

type Application struct {
	TwitchClient *twitch.Client
	Logger       *zap.SugaredLogger
}

func main() {
	cfg := config.New()

	tc := twitch.NewClient(cfg.TwitchUsername, cfg.TwitchOauth)
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	common.StartTime()

	app := &Application{
		TwitchClient: tc,
		Logger:       sugar,
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
		app.Logger.Infow("Successfully connected to Twitch Servers",
			"Bot username", cfg.TwitchUsername,
			"Environment", cfg.Environment,
		)

		common.Send("nourylul", "xd", app.TwitchClient)
	})
	app.TwitchClient.Join("nourylul")

	err := app.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}
}
