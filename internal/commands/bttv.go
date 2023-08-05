package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/lyx0/nourybot/internal/common"
)

func Bttv(target, query string, tc *twitch.Client) {
	reply := fmt.Sprintf("https://betterttv.com/emotes/shared/search?query=%s", query)

	common.Send(target, reply, tc)
}
