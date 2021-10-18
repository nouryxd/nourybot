package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomDog(channel string, client *twitch.Client) {
	reply := api.RandomDog()

	client.Say(channel, reply)
}
