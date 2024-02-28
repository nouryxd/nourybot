package commands

import "fmt"

// Ffz returns the search url for a given query.
func Ffz(query string) string {
	reply := fmt.Sprintf("https://www.frankerfacez.com/emoticons/?q=%s&sort=count-desc&days=0", query)

	return reply
}
