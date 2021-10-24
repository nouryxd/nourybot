package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
)

func SevenTV(target, emote string, nb *bot.Bot) {
	reply := fmt.Sprintf("https://7tv.app/emotes?query=%s", emote)

	nb.Send(target, reply)
}
