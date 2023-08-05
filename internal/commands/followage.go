package commands

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/lyx0/nourybot/internal/commands/decapi"
	"github.com/lyx0/nourybot/internal/common"
	"go.uber.org/zap"
)

// ()currency 10 USD to EUR
func Followage(target, channel, username string, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	resp, err := decapi.Followage(channel, username)
	if err != nil {
		sugar.Error(err)
	}

	common.Send(target, resp, tc)
}
