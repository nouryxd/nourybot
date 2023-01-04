package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/common"
)

func Ffz(target, query string, tc *twitch.Client) {
	reply := fmt.Sprintf("https://www.frankerfacez.com/emoticons/?q=%s", query)

	common.Send(target, reply, tc)
}
