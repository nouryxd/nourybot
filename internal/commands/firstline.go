package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/commands/ivr"
	"github.com/lyx0/nourybot/internal/common"
)

func FirstLine(target, channel, username string, tc *twitch.Client) {
	ivrResponse, err := ivr.FirstLine(channel, username)

	if err != nil {
		common.Send(channel, fmt.Sprint(err), tc)
		return
	}

	common.Send(target, ivrResponse, tc)
}
