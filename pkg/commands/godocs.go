package commands

import (
	"fmt"
	"net/url"
)

// Godocs returns the search link for a given query.
func Godocs(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://godocs.io/?q=%s", query)

	return reply
}
