package commands

import (
	"fmt"
)

// Bttv returns a search string to a specified betterttv search query.
func Bttv(query string) string {
	reply := fmt.Sprintf("https://betterttv.com/emotes/shared/search?query=%s", query)

	return reply
}
