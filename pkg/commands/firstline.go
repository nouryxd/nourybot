package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Firstline(channel, streamer, username string, nb *bot.Bot) {
	ivrResponse, err := ivr.FirstLine(streamer, username)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, ivrResponse)
}
