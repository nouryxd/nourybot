package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/internal/humanize"
)

func Ping(target string, tc *twitch.Client) {
	botUptime := humanize.Time(common.GetUptime())
	commandsUsed := common.GetCommandsUsed()

	reply := fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v", commandsUsed, botUptime)
	common.Send(target, reply, tc)
}
