package commands

import (
	"fmt"
	"net/url"
)

func DuckDuckGo(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://duckduckgo.com/?va=n&hps=1&q=%s", query)

	return reply
}
