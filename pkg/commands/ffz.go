package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

// Ffz responds with a search link for a given ffz emote.
func Ffz(target, emote string, nb *bot.Bot) {
	reply := fmt.Sprintf("https://www.frankerfacez.com/emoticons/?q=%s", emote)

	nb.Send(target, reply)
}
