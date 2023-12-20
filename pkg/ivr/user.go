package ivr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ivrIDByUsernameResponse struct {
	ID string `json:"id"`
}

func IDByUsernameReply(username string) string {
	baseUrl := "https://api.ivr.fi/v2/twitch/user?login="

	resp, err := http.Get(fmt.Sprintf("%s%s", baseUrl, username))
	if err != nil {
		return "xd"
	}

	defer resp.Body.Close()

	responseList := make([]ivrIDByUsernameResponse, 0)
	_ = json.NewDecoder(resp.Body).Decode(&responseList)

	reply := fmt.Sprintf("%s=%s", username, responseList[0].ID)
	return reply
}

func IDByUsername(username string) string {
	baseUrl := "https://api.ivr.fi/v2/twitch/user?login="

	resp, err := http.Get(fmt.Sprintf("%s%s", baseUrl, username))
	if err != nil {
		return "xd"
	}

	defer resp.Body.Close()

	responseList := make([]ivrIDByUsernameResponse, 0)
	_ = json.NewDecoder(resp.Body).Decode(&responseList)

	return responseList[0].ID
}
