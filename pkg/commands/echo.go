package commands

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func Echo(target, message string, tc *twitch.Client) {
	common.Send(target, message, tc)
}
