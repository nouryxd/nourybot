package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Subage(channel string, username string, streamer string, client *twitch.Client) {

	ivrResponse := ivr.Subage(username, streamer)
	client.Say(channel, ivrResponse)
}
