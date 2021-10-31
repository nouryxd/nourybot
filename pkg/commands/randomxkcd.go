package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// RandomXkcd calls the RandomXkcd api and responds with a link to a
// random xkcd comic.
func RandomXkcd(channel string, nb *bot.Bot) {
	reply := api.RandomXkcd()

	nb.Send(channel, reply)
}
