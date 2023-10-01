package main

// This whole file has been pretty much straight up copied
// from: https://github.com/wabarc/go-catbox/blob/main/catbox.go
// since I couldn't otherwise use litterbox instead.

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/wabarc/helper"
)

const (
	//ENDPOINT = "https://catbox.moe/user/api.php"
	ENDPOINT = "https://litterbox.catbox.moe/resources/internals/api.php"
)

type catbox struct {
	Client   *http.Client
	Userhash string
}

func New(client *http.Client) *catbox {
	if client == nil {
		client = &http.Client{
			Timeout: 300 * time.Second,
		}
	}

	return &catbox{
		Client: client,
	}
}

// Upload file or URI to the Catbox. It returns an URL string and error.
func (cat *catbox) Upload(v ...interface{}) (string, error) {
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

func (cat *catbox) rawUpload(b []byte, name string) (string, error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		m.WriteField("userhash", cat.Userhash)
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

func (cat *catbox) fileUpload(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// if size := helper.FileSize(path); size > 209715200 {
	// 	return "", fmt.Errorf("file too large, size: %d MB", size/1024/1024)
	// }

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		//m.WriteField("userhash", cat.Userhash)
		m.WriteField("time", "24h")
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

func (cat *catbox) urlUpload(uri string) (string, error) {
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

func (cat *catbox) Delete(files ...string) error {
	// TODO
	return nil
}

func (app *application) UploadCatbox(target, path string) {
	app.Send(target, "Uploading to catbox... dankCircle")

	cat := New(nil)
	if url, err := cat.fileUpload(path); err != nil {
		app.Send(target, "Something went wrong FeelsBadMan")
	} else {
		app.Send(target, string(url))
	}
}
