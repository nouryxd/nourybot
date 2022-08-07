package decapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

var (
	basePath                 = "https://decapi.me/twitter/latest/"
	twitterUserNotFoundError = "[Error] - [34] Sorry, that page does not exist."
)

func Tweet(username string) (string, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
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
