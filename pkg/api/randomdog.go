package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type randomDogResponse struct {
	Url string `json:"url"`
}

func RandomDog() string {
	response, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject randomDogResponse
	json.Unmarshal(responseData, &responseObject)

	return string(responseObject.Url)
}
