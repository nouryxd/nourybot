package commands

import (
	"fmt"
	"net/url"
)

func Youtube(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", query)

	return reply
}
