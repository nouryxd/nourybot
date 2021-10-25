package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lyx0/nourybot/cmd/bot"
)

func Pyramid(channel, size, emote string, nb *bot.Bot) {
	if size[0] == '.' || size[0] == '/' {
		nb.Send(channel, ":tf:")
		return
	}

	if emote[0] == '.' || emote[0] == '/' {
		nb.Send(channel, ":tf:")
		return
	}

	pyramidSize, err := strconv.Atoi(size)
	pyramidEmote := fmt.Sprint(emote + " ")

	if err != nil {
		nb.Send(channel, "Something went wrong")
	}

	if pyramidSize == 1 {
		nb.Send(channel, fmt.Sprint(pyramidEmote+"..."))
		return
	}

	if pyramidSize > 3 {
		nb.Send(channel, "Max pyramid size is 3")
		return
	}

	for i := 0; i <= pyramidSize; i++ {
		pyramidMessageAsc := strings.Repeat(pyramidEmote, i)
		// fmt.Println(pyramidMessageAsc)

		nb.Send(channel, pyramidMessageAsc)
	}

	for j := pyramidSize - 1; j >= 0; j-- {
		pyramidMessageDesc := strings.Repeat(pyramidEmote, j)
		// fmt.Println(pyramidMessageDesc)

		nb.Send(channel, pyramidMessageDesc)
	}
}
