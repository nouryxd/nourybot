package commands

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api"
	"github.com/sirupsen/logrus"
)

// Currency responds with the conversion rate for two given currencies.
// Example: ()currency 10 usd to eur
func Currency(target, currAmount, currFrom, currTo string, nb *bot.Bot) {
	reply, err := api.Currency(currAmount, currFrom, currTo)
	if err != nil {
		logrus.Info(err)
	}

	nb.Send(target, reply)
}
