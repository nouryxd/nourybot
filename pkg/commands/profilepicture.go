package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func ProfilePicture(channel string, target string, nb *bot.Bot) {
	reply, err := ivr.ProfilePicture(target)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, reply)
}
