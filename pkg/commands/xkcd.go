package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api"
)

func Xkcd(channel string, client *twitch.Client) {
	reply := api.Xkcd()
	client.Say(channel, reply)

}
