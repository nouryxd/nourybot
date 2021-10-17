package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
	log "github.com/sirupsen/logrus"
)

func BttvEmotes(channel string, client *twitch.Client) {
	resp, err := aiden.ApiCall(fmt.Sprintf("api/v1/emotes/%s/bttv", channel))
	if err != nil {
		log.Error(err)
		client.Say(channel, "Something went wrong FeelsBadMan")
	}

	client.Say(channel, string(resp))
}
