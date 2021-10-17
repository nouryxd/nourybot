package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Firstline(channel string, streamer string, username string, client *twitch.Client) {
	ivrResponse, err := ivr.FirstLine(streamer, username)
	if err != nil {
		client.Say(channel, fmt.Sprint(err))
		return
	}
	client.Say(channel, ivrResponse)
}
