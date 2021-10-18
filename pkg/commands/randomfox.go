package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomFox(channel string, client *twitch.Client) {
	reply := api.RandomFox()

	client.Say(channel, reply)
}
