package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func HandleWhisperMessage(whisper twitch.WhisperMessage, client *twitch.Client) {
	log.Info("fn HandleWhisperMessage")
	log.Info(whisper)
	if whisper.Message == "xd" {
		client.Whisper(whisper.User.Name, "xd")
	}
}
