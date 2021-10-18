package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomXkcd(channel string, client *twitch.Client) {
	reply := api.RandomXkcd()
	client.Say(channel, reply)
}
