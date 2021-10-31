package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

// EmoteLookup looks up a given emote and responds with the channel it is
// associated with and its tier.
func EmoteLookup(channel, emote string, nb *bot.Bot) {
	reply, err := ivr.EmoteLookup(emote)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, reply)
}
