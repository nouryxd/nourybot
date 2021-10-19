package handlers

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/sirupsen/logrus"
)

func Command(message twitch.PrivateMessage, nb *bot.Bot) {
	logrus.Info("fn Command")

	// commandName is the actual command name without the prefix.
	commandName := strings.ToLower(strings.SplitN(message.Message, " ", 3)[0][2:])

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
			nb.Send(message.Channel, "xd")
		}
	case "echo":
		if msgLen != 1 {
			nb.Send(message.Channel, cmdParams[1])
			return
		}
	}

}
