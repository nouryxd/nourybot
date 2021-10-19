package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomXkcd(channel string, nb *bot.Bot) {
	reply := api.RandomXkcd()
	nb.Send(channel, reply)
}
