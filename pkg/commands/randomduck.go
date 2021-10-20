package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomDuck(channel string, nb *bot.Bot) {
	reply := api.RandomDuck()

	nb.Send(channel, reply)
}
