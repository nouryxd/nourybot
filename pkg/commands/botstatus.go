package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func BotStatus(channel string, name string, client *twitch.Client) {
	resp, err := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/botStatus/%s?includeLimits=1", name))
	if err != nil {
		client.Say(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	client.Say(channel, resp)

}