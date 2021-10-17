package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func Subage(channel string, username string, streamer string, client *twitch.Client) {

	ivrResponse := api.IvrSubage(username, streamer)
	client.Say(channel, ivrResponse)
}
