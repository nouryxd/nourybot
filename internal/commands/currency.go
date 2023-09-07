package commands

import (
	"fmt"
	"io"
	"net/http"
)

func Currency(currAmount, currFrom, currTo string) (string, error) {
	basePath := "https://decapi.me/misc/currency/"
	from := fmt.Sprintf("?from=%s", currFrom)
	to := fmt.Sprintf("&to=%s", currTo)
	value := fmt.Sprintf("&value=%s", currAmount)

	// https://decapi.me/misc/currency/?from=usd&to=usd&value=10
	resp, err := http.Get(fmt.Sprint(basePath + from + to + value))
	if err != nil {
		return "", ErrInternalServerError
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ErrInternalServerError
	}

	reply := string(body)
	return reply, nil
}
