package commands

import "fmt"

func Ffz(query string) string {
	reply := fmt.Sprintf("https://www.frankerfacez.com/emoticons/?q=%s&sort=count-desc&days=0", query)

	return reply
}
