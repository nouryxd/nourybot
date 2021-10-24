package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type firstLineApiResponse struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Time    string `json:"time"`
	Error   string `json:"error"`
}

// FirstLine returns the first line a given user has sent in a
// given channel.
func FirstLine(channel, username string) (string, error) {
	baseUrl := "https://api.ivr.fi/logs/firstmessage"

	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", baseUrl, channel, username))
	if err != nil {
		log.Error(err)
		return "Something went wrong FeelsBadMan", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject firstLineApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan"), nil
	} else {
		return fmt.Sprintf(username + ": " + responseObject.Message + " (" + responseObject.Time + " ago)."), nil
	}

}
