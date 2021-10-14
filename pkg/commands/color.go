package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func Color(message twitch.PrivateMessage, client *twitch.Client) {
	reply := fmt.Sprintf("@%v, your color is %v", message.User.DisplayName, message.User.Color)

	client.Say(message.Channel, reply)

}
