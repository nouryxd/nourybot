package ivr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// https://api.ivr.fi
type uidApiResponse struct {
	Id    string `json:"id"`
	Error string `json:"error"`
}

func Userid(username string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/resolve/%s", username))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
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
