package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// Xkcd responds with a link to the current xkcd comic and information about i.
func Xkcd(channel string, nb *bot.Bot) {
	reply := api.Xkcd()

	nb.Send(channel, reply)

}
