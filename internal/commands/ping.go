package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/pkg/humanize"
)

func Ping() string {
	botUptime := humanize.Time(common.GetUptime())
	commandsUsed := common.GetCommandsUsed()
	commit := common.GetVersion()

	reply := fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v, Running on commit: %v", commandsUsed, botUptime, commit)

	return reply
}
