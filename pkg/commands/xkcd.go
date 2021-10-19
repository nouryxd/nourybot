package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
)

func Xkcd(channel string, nb *bot.Bot) {
	reply := api.Xkcd()

	nb.Send(channel, reply)

}
