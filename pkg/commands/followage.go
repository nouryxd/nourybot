package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Followage(channel string, streamer string, username string, nb *bot.Bot) {
	ivrResponse, err := ivr.Followage(streamer, username)
	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
	}
	nb.Send(channel, ivrResponse)
}
