package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

// Pingme pings a user.
func Pingme(channel, user string, nb *bot.Bot) {
	response := fmt.Sprintf("🔔 @%s 🔔", user)

	nb.Send(channel, response)
}
