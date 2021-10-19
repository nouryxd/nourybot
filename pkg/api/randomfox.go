package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type randomFoxResponse struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

func RandomFox() string {
	response, err := http.Get("https://randomfox.ca/floof/")
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject randomFoxResponse
	json.Unmarshal(responseData, &responseObject)

	return string(responseObject.Image)
}
