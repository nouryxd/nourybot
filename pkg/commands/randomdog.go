package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomDog(channel string, nb *bot.Bot) {
	reply := api.RandomDog()

	nb.Send(channel, reply)
}
