package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lyx0/nourybot/pkg/common"
)

type xkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}

// Xkcd queries the xkcd api and returns latest xkcd comic.
func Xkcd() (string, error) {
	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return "", ErrInternalServerError
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", ErrInternalServerError
	}
	var responseObject xkcdResponse
	if err = json.Unmarshal(responseData, &responseObject); err != nil {
		return "", ErrInternalServerError
	}

	reply := fmt.Sprint("Current Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	return reply, nil
}

// RandomXkcd queries the xkcd api and returns a random xkcd comic.
func RandomXkcd() (string, error) {
	comicNum := fmt.Sprint(common.GenerateRandomNumber(2899))

	response, err := http.Get(fmt.Sprint("http://xkcd.com/" + comicNum + "/info.0.json"))
	if err != nil {
		return "", ErrInternalServerError
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", ErrInternalServerError
	}
	var responseObject xkcdResponse
	if err = json.Unmarshal(responseData, &responseObject); err != nil {
		return "", ErrInternalServerError
	}

	reply := fmt.Sprint("Random Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	return reply, nil
}
