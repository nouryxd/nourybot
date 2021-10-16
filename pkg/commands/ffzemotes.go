package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func FfzEmotes(channel string, client *twitch.Client) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/emotes/%s/ffz", channel))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	client.Say(channel, string(body))
}
