package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type latestFullResult struct {
	Status  string               `json:"status"`
	Results int                  `json:"results"`
	Data    []latestSingleResult `json:"data"`
}

type latestSingleResult struct {
	ID       string  `json:"id"`
	PreTime  string  `json:"pretime"`
	Release  string  `json:"release"`
	Section  string  `json:"section"`
	Files    string  `json:"files"`
	Size     float64 `json:"size"`
	Status   string  `json:"status"`
	Reason   string  `json:"reason"`
	Group    string  `json:"group"`
	Genre    string  `json:"genre"`
	URL      string  `json:"url"`
	NFO      string  `json:"nfo"`
	NFOImage string  `json:"nfo_image"`
}

func (app *application) PreDBLatest() string {
	baseUrl := "https://api.predb.net/?limit=100"

	resp, err := http.Get(baseUrl)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	var release latestFullResult
	_ = json.Unmarshal([]byte(body), &release)

	var ts []string
	for i := 0; i < release.Results; i++ {
		t := fmt.Sprintf("ID: %v\nRelease timestamp: %v\nRelease: %v\nSection: %v\nFiles: %v\nSize: %v\nStatus: %v\nReason: %v\nRelease group: %v\nRelease genre: %v\npredb.net: %v\nNFO URL: %v\nNFO Image URL: %v\n\n",
			release.Data[i].ID,
			release.Data[i].PreTime,
			release.Data[i].Release,
			release.Data[i].Section,
			release.Data[i].Files,
			release.Data[i].Size,
			release.Data[i].Status,
			release.Data[i].Reason,
			release.Data[i].Group,
			release.Data[i].Genre,
			fmt.Sprint("https://predb.net"+release.Data[i].URL),
			release.Data[i].NFO,
			release.Data[i].NFOImage,
		)
		ts = append(ts, t)

	}

	reply := app.YafUploadString(strings.Join(ts, ""))
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}
	return reply
}

type searchFullResult struct {
	Status  string               `json:"status"`
	Results int                  `json:"results"`
	Data    []searchSingleResult `json:"data"`
}
type searchSingleResult struct {
	ID       int     `json:"id"`
	PreTime  int     `json:"pretime"`
	Release  string  `json:"release"`
	Section  string  `json:"section"`
	Files    int     `json:"files"`
	Size     float64 `json:"size"`
	Status   int     `json:"status"`
	Reason   string  `json:"reason"`
	Group    string  `json:"group"`
	Genre    string  `json:"genre"`
	URL      string  `json:"url"`
	NFO      string  `json:"nfo"`
	NFOImage string  `json:"nfo_image"`
}

func (app *application) PreDBSearch(title string) string {
	escaped := fmt.Sprintf("https://api.predb.net/?q=%s&order_by=pretime&sort=desc&limit=100", url.QueryEscape(title))

	resp, err := http.Get(escaped)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	var release searchFullResult
	_ = json.Unmarshal([]byte(body), &release)

	var ts []string
	for i := 0; i < release.Results; i++ {
		t := fmt.Sprintf("ID: %v\nRelease timestamp: %v\nRelease: %v\nSection: %v\nFiles: %v\nSize: %v\nStatus: %v\nReason: %v\nRelease group: %v\nRelease genre: %v\npredb.net: %v\nNFO URL: %v\nNFO Image URL: %v\n\n",
			release.Data[i].ID,
			release.Data[i].PreTime,
			release.Data[i].Release,
			release.Data[i].Section,
			release.Data[i].Files,
			release.Data[i].Size,
			release.Data[i].Status,
			release.Data[i].Reason,
			release.Data[i].Group,
			release.Data[i].Genre,
			fmt.Sprint("https://predb.net"+release.Data[i].URL),
			release.Data[i].NFO,
			release.Data[i].NFOImage,
		)
		ts = append(ts, t)

	}

	reply := app.YafUploadString(strings.Join(ts, ""))
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}
	return reply
}

func (app *application) PreDBGroup(group string) string {
	escaped := fmt.Sprintf("https://api.predb.net/?group=%s&order_by=pretime&sort=desc&limit=100", url.QueryEscape(group))

	resp, err := http.Get(escaped)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	var release searchFullResult
	_ = json.Unmarshal([]byte(body), &release)

	var ts []string
	for i := 0; i < release.Results; i++ {
		t := fmt.Sprintf("ID: %v\nRelease timestamp: %v\nRelease: %v\nSection: %v\nFiles: %v\nSize: %v\nStatus: %v\nReason: %v\nRelease group: %v\nRelease genre: %v\npredb.net: %v\nNFO URL: %v\nNFO Image URL: %v\n\n",
			release.Data[i].ID,
			release.Data[i].PreTime,
			release.Data[i].Release,
			release.Data[i].Section,
			release.Data[i].Files,
			release.Data[i].Size,
			release.Data[i].Status,
			release.Data[i].Reason,
			release.Data[i].Group,
			release.Data[i].Genre,
			fmt.Sprint("https://predb.net"+release.Data[i].URL),
			release.Data[i].NFO,
			release.Data[i].NFOImage,
		)
		ts = append(ts, t)

	}

	reply := app.YafUploadString(strings.Join(ts, ""))
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return "Something went wrong FeelsBadMan"
	}
	return reply
}
