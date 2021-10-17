package api

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
	Cumulative   Cumulative `json:"cumulative"`
	Streak       SubStreak  `json:"streak"`
	Error        string     `json:"error"`
}

type Cumulative struct {
	Months int `json:"months"`
}

type SubStreak struct {
	Months int `json:"months"`
}

func IvrSubage(username string, streamer string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/subage/%s/%s", username, streamer))
	if err != nil {
		log.Error(err)
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
		return reply
	}

	if responseObject.SubageHidden {

		reply := fmt.Sprintf(username + " has their subscription status hidden. FeelsBadMan")
		return reply
	} else {
		months := fmt.Sprint(responseObject.Cumulative.Months)
		reply := fmt.Sprintf(username + " has been subscribed to " + streamer + " for " + months + " months.")
		return reply
	}
}
