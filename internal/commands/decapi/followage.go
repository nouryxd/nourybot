package decapi

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func Followage(channel, username string) (string, error) {
	var basePath = "https://decapi.me/twitch/followage/"

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	// ?precision is how precise the timestamp should be.
	// precision 4 means:                    1        2         3        4
	// pajlada has been following forsen for 7 years, 4 months, 4 weeks, 1 day
	resp, err := http.Get(fmt.Sprint(basePath + channel + "/" + username + "?precision=4"))
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

	// User tries to look up how long he follows himself.
	if string(body) == followageUserCannotFollowOwn {
		return "You cannot follow yourself.", nil

		// Username is not following the requested channel.
	} else if string(body) == fmt.Sprintf("%s does not follow %s", username, channel) {
		return string(body), nil
	} else {
		reply := fmt.Sprintf("%s has been following %s for %s", username, channel, string(body))
		return reply, nil
	}

}
