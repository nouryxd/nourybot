package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"go.uber.org/zap"
)

func (app *Application) handlePrivateMessage(message twitch.PrivateMessage) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	// roomId is the Twitch UserID of the channel the
	// message originated from.
	roomId := message.Tags["room-id"]

	// If there is no roomId something went wrong.
	if roomId == "" {
		sugar.Errorw("Missing room-id in message tag",
			"roomId", roomId)
		return
	}

	if len(message.Message) >= 2 {
		if message.Message[:2] == "()" {
			//	app.Logger.Infow("handlePrivateMessage",
			//		"message.Message", message.Message)

			handleCommand(message, app.TwitchClient)
			return
		}
	}

	// Message was no command so we just print it.
	app.Logger.Infow("Private Message received",
		"message.Channel", message.Channel,
		"message.User", message.User.DisplayName,
		"message.Message", message.Message)

}
