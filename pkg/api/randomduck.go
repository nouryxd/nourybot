package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type randomDuckResponse struct {
	Url string `json:"url"`
}

// RandomDuck returns a link to a random duck picture.
// API used: https://random-d.uk/api/random
func RandomDuck() string {
	response, err := http.Get("https://random-d.uk/api/random")
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject randomDuckResponse
	json.Unmarshal(responseData, &responseObject)

	return string(responseObject.Url)
}
