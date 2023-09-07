package commands

import (
	"fmt"
)

func Bttv(query string) string {
	reply := fmt.Sprintf("https://betterttv.com/emotes/shared/search?query=%s", query)

	return reply
}
