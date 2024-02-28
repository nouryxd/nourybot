package commands

import (
	"fmt"
	"net/url"
)

// DuckDuckGo returns the search url for a given query.
func DuckDuckGo(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://duckduckgo.com/?va=n&hps=1&q=%s", query)

	return reply
}
