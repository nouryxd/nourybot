package commands

import (
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/humanize"
	"github.com/lyx0/nourybot/pkg/utils"
)

func Ping(channel string, client *twitch.Client, uptime time.Time) {
	commandCount := fmt.Sprint(utils.GetCommandsUsed())
	botUptime := humanize.HumanizeTime(uptime)

	reply := fmt.Sprintf("Pong! :) Commands used: %v, Last restart: %v", commandCount, botUptime)
	client.Say(channel, reply)
}
