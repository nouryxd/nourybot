package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
	"github.com/sirupsen/logrus"
)

func Currency(target string, currAmount string, currFrom string, currTo string, nb *bot.Bot) {
	reply, err := api.Currency(currAmount, currFrom, currTo)
	if err != nil {
		logrus.Info(err)
	}

	nb.Send(target, reply)
}
