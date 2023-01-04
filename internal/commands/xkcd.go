package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/common"
	"go.uber.org/zap"
)

type xkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}

func Xkcd(target string, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		sugar.Error(err)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		sugar.Error(err)
	}
	var responseObject xkcdResponse
	json.Unmarshal(responseData, &responseObject)

	reply := fmt.Sprint("Current Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	common.Send(target, reply, tc)
}

func RandomXkcd(target string, tc *twitch.Client) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	comicNum := fmt.Sprint(common.GenerateRandomNumber(2655))

	response, err := http.Get(fmt.Sprint("http://xkcd.com/" + comicNum + "/info.0.json"))
	if err != nil {
		sugar.Error(err)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		sugar.Error(err)
	}
	var responseObject xkcdResponse
	json.Unmarshal(responseData, &responseObject)

	reply := fmt.Sprint("Random Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	common.Send(target, reply, tc)
}
