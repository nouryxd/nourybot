package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
)

func Weather(channel string, location string, client *twitch.Client) {

	reply := aiden.Weather(fmt.Sprintf("api/v1/misc/weather/%s", location))
	client.Say(channel, reply)

}
