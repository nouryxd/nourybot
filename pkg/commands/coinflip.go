package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/utils"
)

func Coinflip(channel string, nb *bot.Bot) {
	result := utils.GenerateRandomNumber(2)

	if result == 1 {
		nb.Send(channel, "Heads")
	} else {
		nb.Send(channel, "Tails")
	}
}
