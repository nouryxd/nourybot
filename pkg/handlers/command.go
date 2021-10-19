package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/sirupsen/logrus"
)

func Command(message twitch.PrivateMessage, nb *bot.Bot) {
	logrus.Info("fn Command")

	nb.Send("nourybot", "xd")

}
