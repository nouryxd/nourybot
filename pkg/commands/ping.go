package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/utils"
)

func PingCommand(channel string, client *twitch.Client) {
	commandCount := fmt.Sprint(utils.GetCommandsUsed())

	s := fmt.Sprintf("Pong! :) Commands used: %v", commandCount)
	client.Say(channel, s)
}
