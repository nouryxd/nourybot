package commands

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/lyx0/nourybot/internal/commands/decapi"
	"github.com/lyx0/nourybot/internal/common"
	"go.uber.org/zap"
)

func Ffzemotes(target string, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	resp, err := decapi.Ffzemotes(target)
	if err != nil {
		sugar.Error(err)
	}

	common.Send(target, resp, tc)
}
