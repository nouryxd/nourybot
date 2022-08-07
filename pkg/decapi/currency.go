package decapi

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// ()currency 10 USD to EUR
func Currency(currAmount, currFrom, currTo string) (string, error) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	basePath := "https://decapi.me/misc/currency/"
	from := fmt.Sprintf("?from=%s", currFrom)
	to := fmt.Sprintf("&to=%s", currTo)
	value := fmt.Sprintf("&value=%s", currAmount)

	// https://decapi.me/misc/currency/?from=usd&to=usd&value=10
	resp, err := http.Get(fmt.Sprint(basePath + from + to + value))
	if err != nil {
		sugar.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	reply := string(body)
	return reply, nil
}
