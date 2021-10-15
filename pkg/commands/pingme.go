package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func Pingme(channel string, user string, client *twitch.Client) {
	response := fmt.Sprintf("@%s", user)

	client.Say(channel, response)

}
