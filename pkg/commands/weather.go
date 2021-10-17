package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func Weather(channel string, location string, client *twitch.Client) {

	reply, err := aiden.ApiCall(fmt.Sprintf("api/v1/misc/weather/%s", location))
	if err != nil {
		client.Say(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}
	client.Say(channel, reply)

}
