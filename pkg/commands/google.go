package commands

import (
	"fmt"
	"net/url"
)

func Google(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://www.google.com/search?q=%s", query)

	return reply
}
