package main

import "github.com/gempir/go-twitch-irc/v3"

func (app *Application) handleWhisperMessage(message twitch.WhisperMessage) {
	// Print the whisper message for now.
	// TODO: Implement a basic whisper handler.
	app.Logger.Infow("Whisper Message received",
		"message", message,
		"message.User.DisplayName", message.User.DisplayName,
		"message.Message", message.Message)
}
