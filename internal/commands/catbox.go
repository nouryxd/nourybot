// The whole catbox upload functionality has been copied from
// here so that I could use it with litterbox:
// https://github.com/wabarc/go-catbox/blob/main/catbox.go <3
//
// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/wabarc/helper"
	"github.com/wader/goutubedl"
	"go.uber.org/zap"
)

const (
	ENDPOINT = "https://litterbox.catbox.moe/resources/internals/api.php"
)

type Catbox struct {
	Client       *http.Client
	Time         string
	Userhash     string
	TwitchClient *twitch.Client
	Log          *zap.SugaredLogger
}

func NewCat(twitchClient *twitch.Client, log *zap.SugaredLogger) *Catbox {
	client := &http.Client{
		Timeout: 300 * time.Second,
	}

	return &Catbox{
		Client:       client,
		Time:         "24h",
		TwitchClient: twitchClient,
		Log:          log,
	}
}

func DownloadCatbox(target, link string, tc *twitch.Client, log *zap.SugaredLogger) (reply string) {
	cat := NewCat(tc, log)

	go cat.downloadCatbox(target, link)
	return ""
}

func (cat *Catbox) downloadCatbox(target, link string) {
	goutubedl.Path = "yt-dlp"
	var fileName string

	cat.TwitchClient.Say(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		cat.Log.Errorln(err)
	}

	// I don't know why but I need to set it to mp4, otherwise if
	// I use `result.Into.Ext` catbox won't play the video in the
	// browser and say this message:
	// `No video with supported format and MIME type found.`
	rExt := "mp4"

	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		cat.Log.Errorln(err)
	}
	cat.TwitchClient.Say(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	fileName = fmt.Sprintf("%s.%s", uuidFilename, rExt)
	if err != nil {
		cat.Log.Errorln(err)
	}
	f, err := os.Create(fileName)
	cat.TwitchClient.Say(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		cat.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		cat.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()

	if url, err := cat.fileUpload(fileName); err != nil {
		cat.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %v", err))
	} else {
		cat.TwitchClient.Say(target, url)
	}
}

// Upload file or URI to the Catbox. It returns an URL string and error.
func (cat *Catbox) Upload(v ...interface{}) (string, error) {
	if len(v) == 0 {
		return "", fmt.Errorf(`must specify file path or byte slice`)
	}

	switch t := v[0].(type) {
	case string:
		path := t
		parse := func(s string, _ error) (string, error) {
			uri, err := url.Parse(s)
			if err != nil {
				return "", err
			}
			return uri.String(), nil
		}
		switch {
		case helper.IsURL(path):
			return parse(cat.urlUpload(path))
		case helper.Exists(path):
			return parse(cat.fileUpload(path))
		default:
			return "", errors.New(`path invalid`)
		}
	case []byte:
		if len(v) != 2 {
			return "", fmt.Errorf(`must specify file name`)
		}
		return cat.rawUpload(t, v[1].(string))
	}
	return "", fmt.Errorf(`unhandled`)
}

func (cat *Catbox) rawUpload(b []byte, name string) (string, error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		m.WriteField("time", cat.Time)
		//m.WriteField("userhash", cat.Userhash)
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(name))
		if err != nil {
			return
		}
		if _, err = io.Copy(part, bytes.NewBuffer(b)); err != nil {
			return
		}
	}()
	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := cat.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cat *Catbox) fileUpload(path string) (string, error) {
	defer os.Remove(path)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if size := helper.FileSize(path); size > 209715200 {
		return "", fmt.Errorf("file too large, size: %d MB", size/1024/1024)
	}

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		m.WriteField("time", cat.Time)
		m.WriteField("userhash", cat.Userhash)
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(file.Name()))
		if err != nil {
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := cat.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cat *Catbox) urlUpload(uri string) (string, error) {
	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	w.WriteField("reqtype", "urlupload")
	w.WriteField("userhash", cat.Userhash)
	w.WriteField("url", uri)

	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, b)
	req.Header.Add("Content-Type", w.FormDataContentType())

	resp, err := cat.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cat *Catbox) Delete(files ...string) error {
	// TODO
	return nil
}
