package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func BotStatus(channel string, name string, nb *bot.Bot) {
	resp, err := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/botStatus/%s?includeLimits=1", name))
	if err != nil {
		nb.Send(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	nb.Send(channel, resp)
}
