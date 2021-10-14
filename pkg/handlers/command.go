package handlers

import (
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// HandleCommand receives a twitch.PrivateMessage from
// HandlePrivateMessage where it found a command in it.
// HandleCommand passes on the message to the specific
// command handler for further action.
func HandleCommand(message twitch.PrivateMessage, twitchClient *twitch.Client, uptime time.Time) {
	log.Info("fn HandleCommand")

	// Counter that increments on every command call.
	utils.CommandUsed()

	// commandName is the actual command name without the prefix.
	commandName := strings.SplitN(message.Message, " ", 3)[0][2:]

	// cmdParams are additional command inputs.
	// example:
	// weather san antonio
	// is
	// commandName cmdParams[0] cmdParams[1]
	cmdParams := strings.SplitN(message.Message, " ", 500)

	// msgLen is the amount of words in the message without the prefix.
	// Useful for checking if enough cmdParams are given.
	msgLen := len(strings.SplitN(message.Message, " ", -2))

	switch commandName {
	case "":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Why yes, that's my prefix :)")
		}
		return
	case "xd":
		twitchClient.Say(message.Channel, "xd")
		return
	case "echo":
		twitchClient.Say(message.Channel, cmdParams[1])
		return
	case "ping":
		commands.Ping(message.Channel, twitchClient, uptime)
	}
}
