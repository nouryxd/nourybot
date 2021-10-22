package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
	"github.com/sirupsen/logrus"
)

func RandomQuote(channel string, target string, username string, nb *bot.Bot) {
	ivrResponse, err := ivr.RandomQuote(target, username)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	logrus.Info(ivrResponse)
	nb.Send(channel, ivrResponse)
}
