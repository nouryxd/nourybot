package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// RandomCat calls the RandomCat api and responds with a link for a
// random cat image.
func RandomCat(channel string, nb *bot.Bot) {
	reply := api.RandomCat()

	nb.Send(channel, reply)
}
