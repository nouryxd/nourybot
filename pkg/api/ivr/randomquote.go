package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type randomQuoteApiResponse struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Time    string `json:"time"`
	Error   string `json:"Error"`
}

// FirstLine returns the first line a given user has sent in a
// given channel.
func RandomQuote(channel string, username string) (string, error) {
	baseUrl := "https://api.ivr.fi/logs/rq"

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

	var responseObject randomQuoteApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan"), nil
	} else {
		return fmt.Sprintf(username + ": " + responseObject.Message + " (" + responseObject.Time + " ago)."), nil
	}

}
