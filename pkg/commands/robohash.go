package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
)

// Robohash takes the message ID from the message and responds
// with the robohash image link for the message id.
func RoboHash(target string, message twitch.PrivateMessage, nb *bot.Bot) {
	reply := fmt.Sprintf("https://robohash.org/%s", message.ID)

	nb.Send(target, reply)
}
