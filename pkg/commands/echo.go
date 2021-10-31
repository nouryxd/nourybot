package commands

import "github.com/lyx0/nourybot/cmd/bot"

func Echo(channel, message string, nb *bot.Bot) {

	// if message[0] == '.' || message[0] == '/' || message[0] == '!' {
	// 	nb.Send(channel, ":tf:")
	// 	return
	// }

	nb.Send(channel, message)
}
