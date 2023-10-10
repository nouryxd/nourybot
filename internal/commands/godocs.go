package commands

import (
	"fmt"
	"net/url"
)

func Godocs(query string) string {
	query = url.QueryEscape(query)
	reply := fmt.Sprintf("https://godocs.io/?q=%s", query)

	return reply
}
