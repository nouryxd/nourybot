package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
	"github.com/sirupsen/logrus"
)

// commands.Currency(target, cmdParams[1], cmdParams[2], cmdParams[4], nb)
func Currency(target string, currAmount string, currFrom string, currTo string, nb *bot.Bot) {
	reply, err := api.Currency(currAmount, currFrom, currTo)
	if err != nil {
		logrus.Info(err)
	}
	nb.Send(target, reply)
	// logrus.Info(target, reply)
}
