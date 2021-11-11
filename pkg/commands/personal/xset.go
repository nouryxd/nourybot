package personal

import "github.com/lyx0/nourybot/cmd/bot"

// Xset returns the keyboard repeat rate command.
func Xset(target string, nb *bot.Bot) {
	reply := "xset r rate 175 50"

	nb.Send(target, reply)
}
