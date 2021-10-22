package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// https://api.ivr.fi
type followageApiResponse struct {
	User       string `json:"user"`
	UserID     string `json:"userid"`
	Channel    string `json:"channel"`
	ChannelId  string `json:"channelid"`
	FollowedAt string `json:"followedAt"`
	Error      string `json:"error"`
}

// Followage returns the time since a given user followed a given streamer
func Followage(streamer string, username string) (string, error) {
	baseUrl := "https://api.ivr.fi/twitch/subage"

	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", baseUrl, username, streamer))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject followageApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan"), nil
	} else if responseObject.FollowedAt == "" {
		return fmt.Sprintf(username + " is not following " + streamer), nil
	} else {
		d := responseObject.FollowedAt[:10]
		return fmt.Sprintf(username + " has been following " + streamer + " since " + d + "."), nil
	}
}
