package commands

import "github.com/gempir/go-twitch-irc/v2"

func Xd(channel string, client *twitch.Client) {
	client.Say(channel, "xd")
}
