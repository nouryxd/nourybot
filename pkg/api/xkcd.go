package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Xkcd() string {
	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Error(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}
	var responseObject XkcdResponse
	json.Unmarshal(responseData, &responseObject)

	reply := fmt.Sprint("Current Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	return reply
}
