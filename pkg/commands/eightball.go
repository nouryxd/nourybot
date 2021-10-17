package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func EightBall(channel string, client *twitch.Client) {
	resp, err := aiden.ApiCall("api/v1/misc/8ball")
	if err != nil {
		log.Error(err)
		client.Say(channel, "Something went wrong FeelsBadMan")
	}

	client.Say(channel, string(resp))
}
