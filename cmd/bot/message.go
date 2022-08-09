package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"go.uber.org/zap"
)

func (app *Application) handlePrivateMessage(message twitch.PrivateMessage) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	// roomId is the Twitch UserID of the channel the message originated from.
	// If there is no roomId something went really wrong.
	roomId := message.Tags["room-id"]
	if roomId == "" {
		app.Logger.Errorw("Missing room-id in message tag",
			"roomId", roomId,
		)

		return
	}

	// Message was shorter than our prefix is therefore it's irrelevant for us.
	if len(message.Message) >= 2 {
		// Strip the `()` prefix
		if message.Message[:2] == "()" {
			app.handleCommand(message)
			return
		}
	}

	// Message was no command so we just print it.
	app.Logger.Infow("Private Message received",
		// "message", message,
		"message.Channel", message.Channel,
		"message.User", message.User.DisplayName,
		"message.Message", message.Message,
	)

}

func (app *Application) handleWhisperMessage(message twitch.WhisperMessage) {
	// Print the whisper message for now.
	// TODO: Implement a basic whisper handler.
	app.Logger.Infow("Whisper Message received",
		"message", message,
		"message.User.DisplayName", message.User.DisplayName,
		"message.Message", message.Message,
	)
}
