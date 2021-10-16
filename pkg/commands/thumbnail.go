package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/utils"
)

func Thumbnail(channel string, target string, client *twitch.Client) {
	imageHeight := utils.GenerateRandomNumberRange(1040, 1080)
	imageWidth := utils.GenerateRandomNumberRange(1890, 1920)
	response := fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%v-%vx%v.jpg", target, imageWidth, imageHeight)

	client.Say(channel, response)
}
