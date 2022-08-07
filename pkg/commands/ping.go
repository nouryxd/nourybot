package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
	"github.com/lyx0/nourybot/pkg/humanize"
)

func Ping(target string, tc *twitch.Client) {
	botUptime := humanize.Time(common.GetUptime())

	reply := fmt.Sprintf("Pong! :) Last restart: %v", botUptime)
	common.Send(target, reply, tc)
}
