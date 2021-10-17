package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
)

func Weather(channel string, location string, client *twitch.Client) {
	reply := aiden.Weather(location)
	client.Say(channel, reply)

}
