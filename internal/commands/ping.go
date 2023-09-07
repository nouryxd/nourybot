package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/internal/humanize"
)

func Ping() string {
	botUptime := humanize.Time(common.GetUptime())
	commandsUsed := common.GetCommandsUsed()

	reply := fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v", commandsUsed, botUptime)
	return reply
}
