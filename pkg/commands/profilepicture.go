package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func ProfilePicture(channel string, target string, client *twitch.Client) {
	reply, err := ivr.ProfilePicture(target)
	if err != nil {
		client.Say(channel, fmt.Sprint(err))
		return
	}

	client.Say(channel, reply)
}
