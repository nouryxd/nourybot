package commands

import (
	"fmt"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/utils"
)

// Thumbnail responds with the current preview image for a given channel.
func Thumbnail(channel, target string, nb *bot.Bot) {
	imageHeight := utils.GenerateRandomNumberRange(1040, 1080)
	imageWidth := utils.GenerateRandomNumberRange(1890, 1920)

	response := fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%v-%vx%v.jpg", target, imageWidth, imageHeight)

	nb.Send(channel, response)
}
