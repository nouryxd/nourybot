package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func HandleCommand(message twitch.PrivateMessage, twitchClient *twitch.Client) {
	log.Info("fn HandleCommand")
	switch message.Message {
	case "xd":
		twitchClient.Say(message.Channel, "xd")
	}
}
