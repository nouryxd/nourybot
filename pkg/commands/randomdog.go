package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// RandomDog calls the RandomDog api and responds with a link for a
// random dog image.
func RandomDog(channel string, nb *bot.Bot) {
	reply := api.RandomDog()

	nb.Send(channel, reply)
}
