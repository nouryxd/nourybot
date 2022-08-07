package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func Preview(target, channel string, tc *twitch.Client) {
	imageHeight := common.GenerateRandomNumberRange(1040, 1080)
	imageWidth := common.GenerateRandomNumberRange(1890, 1920)

	reply := fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%v-%vx%v.jpg", channel, imageWidth, imageHeight)
	common.Send(target, reply, tc)
}
