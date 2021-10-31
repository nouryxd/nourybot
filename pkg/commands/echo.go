package commands

import "github.com/lyx0/nourybot/cmd/bot"

// Echo responds with the given message.
func Echo(channel, message string, nb *bot.Bot) {

	nb.Send(channel, message)
}
