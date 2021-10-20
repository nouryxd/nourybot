package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lyx0/nourybot/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type XkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}

// RandomXkcd returns a link to a random Xkcd comic.
func RandomXkcd() string {
	comicNum := fmt.Sprint(utils.GenerateRandomNumber(2468))
	response, err := http.Get(fmt.Sprint("http://xkcd.com/" + comicNum + "/info.0.json"))
	if err != nil {
		log.Error(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}
	var responseObject XkcdResponse
	json.Unmarshal(responseData, &responseObject)

	reply := fmt.Sprint("Random Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	return reply
}
