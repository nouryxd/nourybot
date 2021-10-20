package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func FfzEmotes(channel string, nb *bot.Bot) {
	resp, err := aiden.ApiCall(fmt.Sprintf("api/v1/emotes/%s/ffz", channel))

	if err != nil {
		log.Error(err)
		nb.Send(channel, "Something went wrong FeelsBadMan")
	}

	nb.Send(channel, string(resp))
}
