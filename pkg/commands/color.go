package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
)

// Color responds with a users Twitch username color.
func Color(message twitch.PrivateMessage, nb *bot.Bot) {
	reply := fmt.Sprintf("@%v, your color is %v", message.User.DisplayName, message.User.Color)

	nb.Send(message.Channel, reply)

}
