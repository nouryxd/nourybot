package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func Title(channel string, target string, client *twitch.Client) {
	title, err := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/channel/%s/title", target))
	if err != nil {
		client.Say(channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}
	reply := fmt.Sprintf("%s title is: %s", target, title)
	client.Say(channel, reply)
}
