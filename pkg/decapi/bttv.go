package decapi

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func Bttv(username string) (string, error) {
	var basePath = "https://decapi.me/bttv/emotes/"

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	// https://decapi.me/twitter/latest/forsen?url&no_rts
	// ?url adds the url at the end and &no_rts ignores retweets.
	resp, err := http.Get(fmt.Sprint(basePath + username))
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
