package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type subageApiResponse struct {
	User         string     `json:"user"`
	UserID       string     `json:"userId"`
	Channel      string     `json:"channel"`
	ChannelId    string     `json:"channelid"`
	SubageHidden bool       `json:"hidden"`
	Subscribed   bool       `json:"subscribed"`
	FollowedAt   string     `json:"followedAt"`
	Cumulative   cumulative `json:"cumulative"`
	Streak       subStreak  `json:"streak"`
	Error        string     `json:"error"`
}

type cumulative struct {
	Months int `json:"months"`
}

type subStreak struct {
	Months int `json:"months"`
}

var (
	subageBaseUrl = "https://api.ivr.fi/twitch/subage"
)

func Subage(username string, streamer string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", subageBaseUrl, username, streamer))
	if err != nil {
		log.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject subageApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		reply := fmt.Sprintf(responseObject.Error + " FeelsBadMan")
		// client.Say(channel, fmt.Sprintf(responseObject.Error+" FeelsBadMan"))
		return reply, nil
	}

	if responseObject.SubageHidden {

		reply := fmt.Sprintf(username + " has their subscription status hidden. FeelsBadMan")
		return reply, nil
	} else {
		months := fmt.Sprint(responseObject.Cumulative.Months)
		reply := fmt.Sprintf(username + " has been subscribed to " + streamer + " for " + months + " months.")
		return reply, nil
	}
}
