package commands

import (
	"fmt"
	"net/url"

	"github.com/lyx0/nourybot/cmd/bot"
)

func Osrs(target, term string, nb *bot.Bot) {
	reply := fmt.Sprint("https://oldschool.runescape.wiki/?search=" + url.QueryEscape(term))

	nb.Send(target, reply)
}
