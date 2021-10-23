package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

func EmoteLookup(channel string, emote string, nb *bot.Bot) {
	reply, err := ivr.EmoteLookup(emote)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, reply)
}
