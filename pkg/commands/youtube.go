package commands

import (
	"fmt"
	"net/url"
)

// Youtube returns a search link to a given query.
func Youtube(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", query)

	return reply
}
