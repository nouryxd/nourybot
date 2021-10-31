package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

// SevenTV responds with a link to the 7tv search for a given emote.
func SevenTV(target, emote string, nb *bot.Bot) {
	reply := fmt.Sprintf("https://7tv.app/emotes?query=%s", emote)

	nb.Send(target, reply)
}
