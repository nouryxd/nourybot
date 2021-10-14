package commands

import (
	"io/ioutil"
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func EightBall(message twitch.PrivateMessage, client *twitch.Client) {
	resp, err := http.Get("https://customapi.aidenwallis.co.uk/api/v1/misc/8ball")
	if err != nil {
		client.Say(message.Channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		client.Say(message.Channel, "Something went wrong FeelsBadMan")
		log.Error(err)
	}

	client.Say(message.Channel, string(body))

}
