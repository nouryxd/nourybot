package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func Weather(channel, location string, nb *bot.Bot) {
	reply, err := aiden.ApiCall(fmt.Sprintf("api/v1/misc/weather/%s", location))

	if err != nil {
		nb.Send(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
		return
	}

	nb.Send(channel, reply)
}
