package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
)

func RoboHash(target string, message twitch.PrivateMessage, nb *bot.Bot) {
	reply := fmt.Sprintf("https://robohash.org/%s", message.ID)
	nb.Send(target, reply)

}
