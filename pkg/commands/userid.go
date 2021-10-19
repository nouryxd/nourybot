package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Userid(channel string, target string, nb *bot.Bot) {
	reply := ivr.Userid(target)
	nb.Send(channel, reply)
}
