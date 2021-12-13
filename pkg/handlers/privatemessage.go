package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
)

// PrivateMessage checks messages for correctness and forwards
// commands to the command handler.
func PrivateMessage(message twitch.PrivateMessage, nb *bot.Bot) {
	// log.Info("fn PrivateMessage")
	// log.Info(message.Raw)

	// roomId is the Twitch UserID of the channel the message
	// was sent in.
	roomId := message.Tags["room-id"]

	// The message has no room-id so something went wrong.
	if roomId == "" {
		log.Error("Missing room-id in message tag", roomId)
		return
	}

	// General message logging. Not in use currently.
	// db.InsertMessage(nb, message.User.Name, message.Channel, message.User.ID, message.Message)

	// Thing for #pajlada
	if message.Channel == "pajlada" && message.Message == "pajaS 🚨 ALERT" && message.User.Name == "pajbot" && message.Action {
		// log.Info(message.Message)
		nb.SkipChecking("pajlada", "/me PAJAS 🚨 pajaAAAAAAAAAAA")
		return
	}

	// Since our command prefix is () ignore every message
	// that is less than 2
	if len(message.Message) >= 2 {

		// Message starts with (), pass it on to
		// the command handler.
		if message.Message[:2] == "()" {
			Command(message, nb)
			return
		}
	}

	// Message was no command
	log.Infof("[#%s]:%s: %s", message.Channel, message.User.DisplayName, message.Message)

}
