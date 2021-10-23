package commands

import "github.com/lyx0/nourybot/cmd/bot"

func CommandsList(target string, nb *bot.Bot) {
	reply := "https://gist.github.com/lyx0/6e326451ed9602157fcf6b5e40fb1dfd"

	nb.Send(reply, target)
}
