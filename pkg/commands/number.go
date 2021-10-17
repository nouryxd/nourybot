package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func RandomNumber(channel string, client *twitch.Client) {
	reply := api.RandomNumber()
	client.Say(channel, string(reply))
}

func Number(channel string, number string, client *twitch.Client) {
	reply := api.Number(number)
	client.Say(channel, string(reply))
}
