// The whole catbox upload functionality has been copied from
// here so that I could use it with litterbox:
// https://github.com/wabarc/go-catbox/blob/main/catbox.go <3
//
// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

const (
	CATBOX_ENDPOINT = "https://litterbox.catbox.moe/resources/internals/api.php"
	GOFILE_ENDPOINT = "https://store1.gofile.io/uploadFile"
	KAPPA_ENDPOINT  = "https://kappa.lol/api/upload"
	YAF_ENDPOINT    = "https://i.yaf.li/upload"
)

func (app *application) NewUpload(destination, fileName, target, identifier string, msg twitch.PrivateMessage) {

	switch destination {
	case "catbox":
		go app.CatboxUpload(target, fileName, identifier, msg)
	case "yaf":
		go app.YafUpload(target, fileName, identifier, msg)
	case "kappa":
		go app.KappaUpload(target, fileName, identifier, msg)
	case "gofile":
		go app.GofileUpload(target, fileName, identifier, msg)

	}
}

func (app *application) CatboxUpload(target, fileName, identifier string, msg twitch.PrivateMessage) {
	defer os.Remove(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer file.Close()
	app.Send(target, "Uploading to catbox.moe... dankCircle", msg)

	// if size := helper.FileSize(fileName); size > 209715200 {
	// 	return "", fmt.Errorf("file too large, size: %d MB", size/1024/1024)
	// }

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		err := m.WriteField("reqtype", "fileupload")
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			return
		}
		err = m.WriteField("time", "24h")
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			return
		}
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(file.Name()))
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, _ := http.NewRequest(http.MethodPost, CATBOX_ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	client := &http.Client{
		Timeout: 300 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	reply := string(body)
	go app.Models.Uploads.UpdateUploadURL(identifier, reply)
	app.Send(target, fmt.Sprintf("Removing file: %s", fileName), msg)
	app.Send(target, reply, msg)
}

func (app *application) GofileUpload(target, path, identifier string, msg twitch.PrivateMessage) {
	defer os.Remove(path)
	app.Send(target, "Uploading to gofile.io... dankCircle", msg)
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	type gofileData struct {
		DownloadPage string `json:"downloadPage"`
		Code         string `json:"code"`
		ParentFolder string `json:"parentFolder"`
		FileId       string `json:"fileId"`
		FileName     string `json:"fileName"`
		Md5          string `json:"md5"`
	}

	type gofileResponse struct {
		Status string `json:"status"`
		Data   gofileData
	}

	go func() {
		defer pw.Close()

		file, err := os.Open(path) // path to image file
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, GOFILE_ENDPOINT, pr)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{
		Timeout: 300 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		app.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	app.Send(target, "Uploaded PogChamp", msg)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		app.Log.Errorln("Error while reading response:", err)
		return
	}

	jsonResponse := new(gofileResponse)
	if err := json.Unmarshal(body, jsonResponse); err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		app.Log.Errorln("Error while unmarshalling JSON response:", err)
		return
	}

	var reply = jsonResponse.Data.DownloadPage

	go app.Models.Uploads.UpdateUploadURL(identifier, reply)
	app.Send(target, fmt.Sprintf("Removing file: %s", path), msg)
	app.Send(target, reply, msg)
}

func (app *application) KappaUpload(target, path, identifier string, msg twitch.PrivateMessage) {
	defer os.Remove(path)
	app.Send(target, "Uploading to kappa.lol... dankCircle", msg)
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	type kappaResponse struct {
		Link string `json:"link"`
	}

	go func() {
		defer pw.Close()

		err := form.WriteField("name", "xd")
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		file, err := os.Open(path) // path to image file
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, KAPPA_ENDPOINT, pr)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{
		Timeout: 300 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		app.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	app.Send(target, "Uploaded PogChamp", msg)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		app.Log.Errorln("Error while reading response:", err)
		return
	}

	jsonResponse := new(kappaResponse)
	if err := json.Unmarshal(body, jsonResponse); err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		app.Log.Errorln("Error while unmarshalling JSON response:", err)
		return
	}

	var reply = jsonResponse.Link

	go app.Models.Uploads.UpdateUploadURL(identifier, reply)
	app.Send(target, fmt.Sprintf("Removing file: %s", path), msg)
	app.Send(target, reply, msg)
}

func (app *application) YafUpload(target, path, identifier string, msg twitch.PrivateMessage) {
	defer os.Remove(path)
	app.Send(target, "Uploading to yaf.li... dankCircle", msg)
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()

		err := form.WriteField("name", "xd")
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		file, err := os.Open(path) // path to image file
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
			os.Remove(path)
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, YAF_ENDPOINT, pr)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{
		Timeout: 300 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		app.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		os.Remove(path)
		app.Log.Errorln("Error while reading response:", err)
		return
	}

	var reply = string(body[:])

	go app.Models.Uploads.UpdateUploadURL(identifier, reply)
	app.Send(target, fmt.Sprintf("Removing file: %s", path), msg)
	app.Send(target, reply, msg)
}
