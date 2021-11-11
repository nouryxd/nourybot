package personal

import "github.com/lyx0/nourybot/cmd/bot"

// Justinfan responds with the Jusinfan user chatterino has used before.
func Justinfan(target string, nb *bot.Bot) {
	reply := "64537"

	nb.Send(target, reply)
}
