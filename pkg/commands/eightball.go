package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func EightBall(channel string, nb *bot.Bot) {
	resp, err := aiden.ApiCall("api/v1/misc/8ball")

	if err != nil {
		log.Error(err)
		nb.Send(channel, "Something went wrong FeelsBadMan")
	}

	nb.Send(channel, string(resp))
}
