package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

// RandomNumber calls the numbers api with a random number.
func RandomNumber(channel string, nb *bot.Bot) {
	reply := api.RandomNumber()

	nb.Send(channel, string(reply))
}

// Number calls the numbers api with a given number.
func Number(channel, number string, nb *bot.Bot) {
	reply := api.Number(number)

	nb.Send(channel, string(reply))
}
