package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomCat(channel string, nb *bot.Bot) {
	reply := api.RandomCat()

	nb.Send(channel, reply)
}
