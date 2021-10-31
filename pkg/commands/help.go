package commands

import "github.com/lyx0/nourybot/cmd/bot"

// Help responds with basic information about the bot.
func Help(target string, nb *bot.Bot) {
	reply := "Bot made in Go by @Nouryqt, Prefix: (), Commands: https://gist.github.com/lyx0/6e326451ed9602157fcf6b5e40fb1dfd"

	nb.Send(target, reply)
}
