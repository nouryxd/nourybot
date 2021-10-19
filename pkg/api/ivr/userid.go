package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// https://api.ivr.fi
type uidApiResponse struct {
	Id    string `json:"id"`
	Error string `json:"error"`
}

func Userid(username string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/resolve/%s", username))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject uidApiResponse
	json.Unmarshal(body, &responseObject)

	// User not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan")
	} else {
		return fmt.Sprintf(responseObject.Id)
	}
}
