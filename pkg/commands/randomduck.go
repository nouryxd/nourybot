package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// RandomDuck calls the RandomDuck api and responds with a link for a
// random duck image.
func RandomDuck(channel string, nb *bot.Bot) {
	reply := api.RandomDuck()

	nb.Send(channel, reply)
}
