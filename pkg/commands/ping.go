package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
	"github.com/lyx0/nourybot/pkg/humanize"
)

func Ping(target string, tc *twitch.Client) {
	botUptime := humanize.Time(common.GetUptime())
	commandsUsed := common.GetCommandsUsed()

	reply := fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v", commandsUsed, botUptime)
	common.Send(target, reply, tc)
}
