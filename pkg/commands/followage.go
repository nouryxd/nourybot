package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

// Followage responds with a given users time since he followed a streamer.
func Followage(channel, streamer, username string, nb *bot.Bot) {
	ivrResponse, err := ivr.Followage(streamer, username)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, ivrResponse)
}
