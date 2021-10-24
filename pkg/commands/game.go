package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func Game(channel, name string, nb *bot.Bot) {
	game, err := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/channel/%s/game", name))

	if err != nil {
		nb.Send(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	reply := fmt.Sprintf("@%s was last seen playing: %s", name, game)
	nb.Send(channel, reply)
}
