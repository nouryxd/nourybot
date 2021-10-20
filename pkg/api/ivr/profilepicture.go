package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// https://api.ivr.fi
type pfpApiResponse struct {
	Logo  string `json:"logo"`
	Error string `json:"error"`
}

var (
	baseUrl = "https://api.ivr.fi/twitch/resolve"
)

// ProfilePicture returns a link to a given users profilepicture.
func ProfilePicture(username string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, username))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject pfpApiResponse
	json.Unmarshal(body, &responseObject)

	// User not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan"), nil
	} else {
		return fmt.Sprintf(responseObject.Logo), nil
	}
}
