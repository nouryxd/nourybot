package commands

import "github.com/lyx0/nourybot/cmd/bot"

func Justinfan(target string, nb *bot.Bot) {
	reply := "64537"

	nb.Send(target, reply)
}
