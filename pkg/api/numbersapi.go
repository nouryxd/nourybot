package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RandomNumber returns a string containg fun facts about a random number.
// API used: http://numbersapi.com
func RandomNumber() string {
	response, err := http.Get("http://numbersapi.com/random/trivia")
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	return string(responseData)
}

// Number returns a string containing fun facts about a given number.
// API used: http://numbersapi.com
func Number(number string) string {
	response, err := http.Get(fmt.Sprint("http://numbersapi.com/" + string(number)))

	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	return string(responseData)
}
