package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func Uptime(channel string, name string, nb *bot.Bot) {
	resp, err := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/channel/%s/uptime", name))
	if err != nil {
		nb.Send(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	nb.Send(channel, resp)

}
