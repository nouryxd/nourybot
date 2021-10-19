package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type randomCatResponse struct {
	File string `json:"file"`
}

func RandomCat() string {
	response, err := http.Get("https://aws.random.cat/meow")
	if err != nil {
		log.Error(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject randomCatResponse
	json.Unmarshal(responseData, &responseObject)

	return string(responseObject.File)
}
