package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Subage(channel, username, streamer string, nb *bot.Bot) {
	ivrResponse, err := ivr.Subage(username, streamer)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, ivrResponse)
}
