package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// RandomFox calls the RandomFox api and responds with a link for a
// random fox image.
func RandomFox(channel string, nb *bot.Bot) {
	reply := api.RandomFox()

	nb.Send(channel, reply)
}
