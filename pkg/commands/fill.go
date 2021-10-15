package commands

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

func Fill(channel string, emote string, client *twitch.Client) {
	if emote[0] == '.' || emote[0] == '/' {
		client.Say(channel, ":tf:")
		return
	}

	// Get the length of the emote
	emoteLength := (len(emote) + 1)

	// Check how often the emote fits into a single message
	repeatCount := (499 / emoteLength)

	reply := strings.Repeat(fmt.Sprintf(emote+" "), repeatCount)
	client.Say(channel, reply)

}
