package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomCat(channel string, client *twitch.Client) {
	reply := api.RandomCat()

	client.Say(channel, reply)
}
