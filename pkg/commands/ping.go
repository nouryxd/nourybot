package commands

import (
	"fmt"

	"github.com/nouryxd/nourybot/pkg/common"
	"github.com/nouryxd/nourybot/pkg/humanize"
)

// Ping returns information about the bot.
func Ping(env string) string {
	botUptime := humanize.Time(common.GetUptime())
	commandsUsed := common.GetCommandsUsed()
	commit := common.GetCommit()
	goVersion := common.GetGoVersion()

	return fmt.Sprintf("Pong! :) Last restart: %v, Commands used: %v, Go version: %v, Commit: %v, Env: %v", botUptime, commandsUsed, goVersion, commit, env)
}
