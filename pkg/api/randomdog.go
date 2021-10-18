package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type randomDogResponse struct {
	Url string `json:"url"`
}

func RandomDog() string {
	response, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		log.Fatalln(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject randomDogResponse
	json.Unmarshal(responseData, &responseObject)

	return string(responseObject.Url)
}
