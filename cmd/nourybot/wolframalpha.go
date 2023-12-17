package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (app *application) WolframAlphaQuery(query string) string {
	escaped := url.QueryEscape(query)
	url := fmt.Sprintf("http://api.wolframalpha.com/v1/result?appid=%s&i=%s", app.Config.wolframAlphaAppID, escaped)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	reply := string(body)
	return reply

}
