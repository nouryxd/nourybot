package commands

import "github.com/lyx0/nourybot/cmd/bot"

func Echo(channel, message string, nb *bot.Bot) {
	nb.Send(channel, message)
}
