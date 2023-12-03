package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// uploadPaste uploads a given text to a pastebin site and returns the link
//
// this whole function was pretty much yoinked from here
// https://github.com/zneix/haste-client/blob/master/main.go <3
func (app *application) uploadPaste(text string) (string, error) {
	const hasteURL = "https://haste.dank.pw"
	const apiRoute = "/documents"
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	type pasteResponse struct {
		Key string `json:"key,omitempty"`
	}

	req, err := http.NewRequest("POST", hasteURL+apiRoute, bytes.NewBufferString(text))
	if err != nil {
		app.Log.Errorln("Could not upload paste:", err)
		return "", err
	}

	req.Header.Set("User-Agent", "nourybot")

	resp, err := httpClient.Do(req)
	if err != nil {
		app.Log.Errorln("Error while sending HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		app.Log.Errorln("Failed to upload data, server responded with", resp.StatusCode)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Log.Errorln("Error while reading response:", err)
		return "", err
	}

	jsonResponse := new(pasteResponse)
	if err := json.Unmarshal(body, jsonResponse); err != nil {
		app.Log.Errorln("Error while unmarshalling JSON response:", err)
		return "", err
	}

	finalURL := hasteURL + "/" + jsonResponse.Key

	return finalURL, nil
}
