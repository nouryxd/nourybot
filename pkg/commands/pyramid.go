package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

func Pyramid(channel string, size string, emote string, client *twitch.Client) {
	if size[0] == '.' || size[0] == '/' {
		client.Say(channel, ":tf:")
		return
	}

	if emote[0] == '.' || emote[0] == '/' {
		client.Say(channel, ":tf:")
		return
	}

	pyramidSize, err := strconv.Atoi(size)
	pyramidEmote := fmt.Sprint(emote + " ")

	if err != nil {
		client.Say(channel, "Something went wrong")
	}

	if pyramidSize == 1 {
		client.Say(channel, fmt.Sprint(pyramidEmote+"..."))
		return
	}

	if pyramidSize > 20 {
		client.Say(channel, "Max pyramid size is 20")
		return
	}

	for i := 0; i <= pyramidSize; i++ {
		pyramidMessageAsc := strings.Repeat(pyramidEmote, i)
		// fmt.Println(pyramidMessageAsc)
		client.Say(channel, pyramidMessageAsc)
	}
	for j := pyramidSize - 1; j >= 0; j-- {
		pyramidMessageDesc := strings.Repeat(pyramidEmote, j)
		// fmt.Println(pyramidMessageDesc)
		client.Say(channel, pyramidMessageDesc)
	}
}
