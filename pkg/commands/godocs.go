package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

func Godocs(channel string, searchTerm string, nb *bot.Bot) {
	resp := fmt.Sprintf("https://godocs.io/?q=%s", searchTerm)

	nb.Send(channel, resp)
}
