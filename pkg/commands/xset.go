package commands

import "github.com/lyx0/nourybot/cmd/bot"

func Xset(target string, nb *bot.Bot) {
	reply := "xset r rate 175 50"

	nb.Send(target, reply)
}
