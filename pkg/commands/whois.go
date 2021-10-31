package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

// Whois responds with information about a given users Twitch account.
func Whois(target, user string, nb *bot.Bot) {
	reply := ivr.Whois(user)

	nb.Send(target, reply)
}
