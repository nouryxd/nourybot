package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func Userid(channel string, target string, client *twitch.Client) {
	reply := ivr.Userid(target)
	client.Say(channel, reply)
}
