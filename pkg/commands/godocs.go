package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

// Godocs responds with a search for a given term on godocs.io
func Godocs(channel, searchTerm string, nb *bot.Bot) {
	resp := fmt.Sprintf("https://godocs.io/?q=%s", searchTerm)

	nb.Send(channel, resp)
}
