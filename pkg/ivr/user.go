package ivr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ivrResponse struct {
	ID     string    `json:"id"`
	Stream ivrStream `json:"stream"`
}

type ivrStream struct {
	Title string  `json:"title"`
	Game  ivrGame `json:"game"`
}

type ivrGame struct {
	DisplayName string `json:"displayName"`
}

// TitleByUsername returns the current title of a supplied twitch channel.
func TitleByUsername(login string) string {
	baseUrl := "https://api.ivr.fi/v2/twitch/user?login="

	resp, err := http.Get(fmt.Sprintf("%s%s", baseUrl, login))
	if err != nil {
		return "xd"
	}

	defer resp.Body.Close()

	responseList := make([]ivrResponse, 0)
	_ = json.NewDecoder(resp.Body).Decode(&responseList)
	if len(responseList) == 0 {
		return "xd"
	}

	return responseList[0].Stream.Title
}

// GameByUsername returns the current game of a supplied twitch channel.
func GameByUsername(login string) string {
	baseUrl := "https://api.ivr.fi/v2/twitch/user?login="

	resp, err := http.Get(fmt.Sprintf("%s%s", baseUrl, login))
	if err != nil {
		return "xd"
	}

	defer resp.Body.Close()

	responseList := make([]ivrResponse, 0)
	_ = json.NewDecoder(resp.Body).Decode(&responseList)
	if len(responseList) == 0 {
		return "xd"
	}

	return responseList[0].Stream.Game.DisplayName
}

// IDByUsernameReply returns the twitch user id of a supplied
// twitch username in the format of "username=id"
func IDByUsernameReply(username string) string {
	baseUrl := "https://api.ivr.fi/v2/twitch/user?login="

	resp, err := http.Get(fmt.Sprintf("%s%s", baseUrl, username))
	if err != nil {
		return "xd"
	}

	responseList := make([]ivrResponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&responseList)
	if len(responseList) == 0 {
		return "xd"
	}

	reply := fmt.Sprintf("%s=%s", username, responseList[0].ID)
	return reply
}

// IDByUsername returns the twitch user id for a given twitch username.
func IDByUsername(username string) string {
	baseUrl := "https://api.ivr.fi/v2/twitch/user?login="

	resp, err := http.Get(fmt.Sprintf("%s%s", baseUrl, username))
	if err != nil {
		return "xd"
	}

	defer resp.Body.Close()

	responseList := make([]ivrResponse, 0)
	_ = json.NewDecoder(resp.Body).Decode(&responseList)
	if len(responseList) == 0 {
		return "xd"
	}

	return responseList[0].ID
}
