package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Followage(channel string, streamer string, username string, client *twitch.Client) {
	ivrResponse, err := ivr.Followage(streamer, username)
	if err != nil {
		client.Say(channel, fmt.Sprint(err))
	}
	client.Say(channel, ivrResponse)
}
