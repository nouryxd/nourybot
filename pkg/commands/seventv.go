package commands

import "fmt"

func SevenTV(query string) string {
	reply := fmt.Sprintf("https://7tv.app/emotes?page=1&query=%s", query)

	return reply
}

