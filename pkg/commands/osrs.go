package commands

import (
	"fmt"
	"net/url"
)

// OSRS returns the search link for a given query.
func OSRS(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://oldschool.runescape.wiki/?search=%s", query)

	return reply
}
