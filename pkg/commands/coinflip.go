package commands

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func Coinflip(target string, tc *twitch.Client) {
	flip := common.GenerateRandomNumber(2)

	switch flip {
	case 0:
		common.Send(target, "Heads!", tc)
		return
	case 1:
		common.Send(target, "Tails!", tc)
		return
	default:
		common.Send(target, "Heads!", tc)
		return
	}

}
