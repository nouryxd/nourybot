package commands

import (
	"fmt"
	"net/url"
)

// Google returns the search link for a given query.
func Google(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://www.google.com/search?q=%s", query)

	return reply
}
