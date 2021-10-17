package aiden

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ApiCall(uri string) string {
	resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/%s", uri))
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
