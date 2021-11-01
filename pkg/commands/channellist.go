package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/db"
)

func ChannelList(target string, nb *bot.Bot) {
	db.ListChannel(nb)
	nb.Send(target, "Whispered you the list of channel.")
}
