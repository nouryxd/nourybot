package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

func Ping(target string, nb *bot.Bot) {

	reply := fmt.Sprintf("Pong! :) Last restart: %v", nb.Uptime)

	nb.Send(target, reply)
}
