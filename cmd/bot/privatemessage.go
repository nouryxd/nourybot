package main

import "github.com/gempir/go-twitch-irc/v3"

func (app *Application) handlePrivateMessage(message twitch.PrivateMessage) {
	// roomId is the Twitch UserID of the channel the
	// message originated from.
	roomId := message.Tags["room-id"]

	// If there is no roomId something went wrong.
	if roomId == "" {
		app.Logger.Error("Missing room-id in message tag ", roomId)
		return
	}

	if len(message.Message) >= 2 {
		if message.Message[:2] == "()" {
			// TODO: Command Handling
			app.handleCommand(message)
			// app.logger.Infof("[Command detected]: ", message.Message)
			return
		}
	}

	// Message was no command so we just print it.
	app.Logger.Infof("[#%s]:%s: %s", message.Channel, message.User.DisplayName, message.Message)

}