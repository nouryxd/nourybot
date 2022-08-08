package decapi

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func Tweet(username string) (string, error) {
	var basePath = "https://decapi.me/twitter/latest/"

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	// https://decapi.me/twitter/latest/forsen?url&no_rts
	// ?url adds the url at the end and &no_rts ignores retweets.
	resp, err := http.Get(fmt.Sprint(basePath + username + "?url" + "&no_rts"))
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

	// If the response was a known error message return a message with the error.
	if string(body) == twitterUserNotFoundError {
		return "Something went wrong: Twitter username not found", err
	} else { // No known error was found, probably a tweet.
		resp := fmt.Sprintf("Latest Tweet from @%s: \"%s \"", username, body)
		return resp, nil
	}
}
