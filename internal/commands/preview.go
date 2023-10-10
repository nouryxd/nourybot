package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/internal/common"
)

func Preview(channel string) string {
	imageHeight := common.GenerateRandomNumberRange(1040, 1080)
	imageWidth := common.GenerateRandomNumberRange(1890, 1920)

	reply := fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%v-%vx%v.jpg", channel, imageWidth, imageHeight)
	return reply
}
