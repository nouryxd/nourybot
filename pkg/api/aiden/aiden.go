package aiden

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	basePath = "https://customapi.aidenwallis.co.uk/"
)

// ApiCall calls https://customapi.aidenwallis.co.uk/ with
// a given uri and returns the result and an error.
func ApiCall(uri string) (string, error) {
	resp, err := http.Get(fmt.Sprint(basePath + uri))
	if err != nil {
		log.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "Something went wrong FeelsBadMan", err
	}
	return string(body), nil
}
