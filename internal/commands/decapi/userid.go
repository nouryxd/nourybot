package decapi

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func GetIdByLogin(login string) (string, error) {
	var basePath = "https://decapi.me/twitch/id/"

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	resp, err := http.Get(fmt.Sprint(basePath + login))
	if err != nil {
		sugar.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	reply := string(body)
	return reply, nil

}
