package commands

import "github.com/gempir/go-twitch-irc/v2"

func Echo(channel string, message string, client *twitch.Client) {
	client.Say(channel, message)
}
