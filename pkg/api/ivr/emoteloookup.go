package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type emoteLookupResponse struct {
	Channel   string `json:"channel"`
	EmoteCode string `json:"emotecode"`
	Tier      string `json:"tier"`
	Emote     string `json:"emote"`
	Error     string `json:"error"`
}

// ProfilePicture returns a link to a given users profilepicture.
func EmoteLookup(emote string) (string, error) {
	baseUrl := "https://api.ivr.fi/twitch/emotes"

	resp, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, emote))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject emoteLookupResponse
	json.Unmarshal(body, &responseObject)

	// log.Info(responseObject.Tier)

	// Emote not found
	if responseObject.Error != "" {
		return fmt.Sprintf(responseObject.Error + " FeelsBadMan"), nil
	} else {
		return fmt.Sprintf("%s is a Tier %v emote to channel %s.", responseObject.EmoteCode, responseObject.Tier, responseObject.Channel), nil
	}
}
