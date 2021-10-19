package commands

import (
	"fmt"
	"strings"

	"github.com/lyx0/nourybot/cmd/bot"
)

func Fill(channel string, emote string, nb *bot.Bot) {
	if emote[0] == '.' || emote[0] == '/' {
		nb.Send(channel, ":tf:")
		return
	}

	// Get the length of the emote
	emoteLength := (len(emote) + 1)

	// Check how often the emote fits into a single message
	repeatCount := (499 / emoteLength)

	reply := strings.Repeat(fmt.Sprintf(emote+" "), repeatCount)
	nb.Send(channel, reply)

}
