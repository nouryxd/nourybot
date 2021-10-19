package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/humanize"
	"github.com/lyx0/nourybot/pkg/utils"
)

func Ping(target string, nb *bot.Bot) {
	commandCount := fmt.Sprint(utils.GetCommandsUsed())
	botUptime := humanize.Time(nb.Uptime)

	reply := fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v", commandCount, botUptime)

	nb.Send(target, reply)
}
