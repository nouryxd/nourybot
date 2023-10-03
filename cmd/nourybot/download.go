package main

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wader/goutubedl"
)

func (app *application) Download(target, link string) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
	}
	rExt := "mp4"
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
	}
	app.Send(target, "Downloaded..")
	fn, err := uuid.NewUUID()
	if err != nil {
		app.Log.Errorln(err)
	}
	f, err := os.Create(fmt.Sprintf("%s.%s", fn, rExt))
	app.Send(target, fmt.Sprintf("Filename: %s.%s", fn, rExt))

	if err != nil {
		app.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()
	duration := 5 * time.Second
	app.Send(target, "ResidentSleeper ..")
	time.Sleep(duration)

	app.upload(target, fmt.Sprintf("%s.%s", fn, rExt))

}

func (app *application) upload(target, path string) {
	const URL = "https://i.yaf.ee/upload"
	app.Send(target, "Uploading .. dankCircle")
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()

		err := form.WriteField("name", "xd")
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		file, err := os.Open(path) // path to image file
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, URL, pr)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		app.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	app.Send(target, "Uploaded .. PogChamp")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		app.Log.Errorln("Error while reading response:", err)
		return
	}

	var reply = string(body[:])

	app.Send(target, fmt.Sprintf("Removing file: %s", path))
	os.Remove(path)
	app.Send(target, reply)
}
