package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func Game(channel string, name string, client *twitch.Client) {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/twitch/channel/%s/game", name))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	client.Say(channel, fmt.Sprintf("%s is playing: %s", name, string(body)))
}
