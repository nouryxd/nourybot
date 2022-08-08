package commands

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/commands/decapi"
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

// ()currency 10 USD to EUR
func Currency(target, currAmount, currFrom, currTo string, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	resp, err := decapi.Currency(currAmount, currFrom, currTo)
	if err != nil {
		sugar.Error(err)
	}

	common.Send(target, resp, tc)
}
