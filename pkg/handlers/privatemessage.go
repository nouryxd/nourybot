package handlers

import (
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/config"
	log "github.com/sirupsen/logrus"
)

// HandlePrivateMessage takes in a twitch.Privatemessage,
// *twitch.Client and *config.Config and has the logic to decide if the provided
// PrivateMessage is a command or not and passes it on accordingly.
// Typical twitch message tags https://paste.ivr.fi/nopiradodo.lua
func HandlePrivateMessage(message twitch.PrivateMessage, client *twitch.Client, cfg *config.Config, uptime time.Time) {
	log.Info("fn HandlePrivateMessage")
	// log.Info(message)

	// roomId is the Twitch UserID of the channel the message
	// was sent in.
	roomId := message.Tags["room-id"]

	// The message has no room-id so something went wrong.
	if roomId == "" {
		log.Errorf("Missing room-id in message tag", roomId)
		return
	}

	// Message was sent from the Bot. Don't act on it
	// so that we don't repeat ourself.
	botUserId := cfg.BotUserId
	if message.Tags["user-id"] == botUserId {
		return
	}

	// Since our command prefix is () ignore every message
	// that is less than 2
	if len(message.Message) >= 2 {

		// Message starts with (), pass it on to
		// the command handler.
		if message.Message[:2] == "()" {
			HandleCommand(message, client, uptime)
			return
		}
	}

	// Message was no command
	// log.Info(message)

}
