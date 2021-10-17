package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/api/aiden"
)

func Game(channel string, name string, client *twitch.Client) {

	game := aiden.ApiCall(fmt.Sprintf("api/v1/twitch/channel/%s/game", name))
	reply := fmt.Sprintf("@%s was last seen playing: %s", name, game)
	client.Say(channel, reply)
}
