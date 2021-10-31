package commands

import "github.com/lyx0/nourybot/cmd/bot"

// Xd responds with a xd
func Xd(channel string, nb *bot.Bot) {
	nb.Send(channel, "xd")
}
