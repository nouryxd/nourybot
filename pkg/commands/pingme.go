package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

func Pingme(channel, user string, nb *bot.Bot) {
	response := fmt.Sprintf("@%s", user)

	nb.Send(channel, response)
}
