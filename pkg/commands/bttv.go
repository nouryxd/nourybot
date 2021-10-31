package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

// Bttv responds with a search for a given bttv emote.
func Bttv(target, emote string, nb *bot.Bot) {
	reply := fmt.Sprintf("https://betterttv.com/emotes/shared/search?query=%s", emote)

	nb.Send(target, reply)
}
