package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomFox(channel string, nb *bot.Bot) {
	reply := api.RandomFox()

	nb.Send(channel, reply)
}
