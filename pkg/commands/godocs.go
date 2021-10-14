package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func Godocs(channel string, searchTerm string, client *twitch.Client) {
	resp := fmt.Sprintf("https://godocs.io/?q=%s", searchTerm)

	client.Say(channel, resp)
}
