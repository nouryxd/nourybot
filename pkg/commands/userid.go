package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

// Userid responds with the userid for a given user.
func Userid(channel, target string, nb *bot.Bot) {
	reply := ivr.Userid(target)

	nb.Send(channel, reply)
}
