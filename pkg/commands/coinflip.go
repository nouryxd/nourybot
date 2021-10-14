package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/utils"
)

func Coinflip(channel string, client *twitch.Client) {
	result := utils.GenerateRandomNumber(2)

	if result == 1 {
		client.Say(channel, "Heads")
	} else {
		client.Say(channel, "Tails")
	}
}
