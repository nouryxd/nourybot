package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

// Firstline responds a given users first message in a given channel.
func Firstline(channel, streamer, username string, nb *bot.Bot) {
	ivrResponse, err := ivr.FirstLine(streamer, username)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, ivrResponse)
}
