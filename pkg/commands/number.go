package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomNumber(channel string, nb *bot.Bot) {
	reply := api.RandomNumber()
	nb.Send(channel, string(reply))
}

func Number(channel string, number string, nb *bot.Bot) {
	reply := api.Number(number)
	nb.Send(channel, string(reply))
}
