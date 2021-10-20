package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func Title(channel string, target string, nb *bot.Bot) {
	title, err := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/channel/%s/title", target))

	if err != nil {
		nb.Send(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	reply := fmt.Sprintf("%s title is: %s", target, title)
	nb.Send(channel, reply)
}
