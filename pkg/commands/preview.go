package commands

import (
	"fmt"

	"github.com/nouryxd/nourybot/pkg/common"
)

// Preview returns a link to an almost live image of a given twitch stream
// if the channel is currently live.
func Preview(channel string) string {
	imageHeight := common.GenerateRandomNumberRange(1040, 1080)
	imageWidth := common.GenerateRandomNumberRange(1890, 1920)

	reply := fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%v-%vx%v.jpg", channel, imageWidth, imageHeight)
	return reply
}
