package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/api/ivr"
)

// RandomQuote calls the RandomQuote api and responds with a link for a
// random quote image.
func RandomQuote(channel, target, username string, nb *bot.Bot) {
	ivrResponse, err := ivr.RandomQuote(target, username)

	if err != nil {
		nb.Send(channel, fmt.Sprint(err))
		return
	}

	nb.Send(channel, ivrResponse)
}
