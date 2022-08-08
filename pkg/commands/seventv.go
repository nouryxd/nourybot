package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func Seventv(target, emote string, tc *twitch.Client) {
	reply := fmt.Sprintf("https://7tv.app/emotes?query=%s", emote)

	common.Send(target, reply, tc)
}
