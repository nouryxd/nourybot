package aiden

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Weather(location string) string {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/misc/weather/%s", location))
	if err != nil {
		log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "Something went wrong FeelsBadMan"
	}
	return string(body)
}
