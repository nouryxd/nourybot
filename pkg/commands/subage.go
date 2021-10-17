package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Subage(channel string, username string, streamer string, client *twitch.Client) {

	ivrResponse, err := ivr.Subage(username, streamer)
	if err != nil {
		client.Say(channel, fmt.Sprint(err))
		return
	}
	client.Say(channel, ivrResponse)
}
