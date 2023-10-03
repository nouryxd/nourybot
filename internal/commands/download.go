package commands

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
	"github.com/wader/goutubedl"
	"go.uber.org/zap"
)

type downloader struct {
	twitchClient *twitch.Client
	Log          *zap.SugaredLogger
}

func Download(target, link string, tc *twitch.Client, log *zap.SugaredLogger) (reply string) {
	dloader := &downloader{
		Log:          log,
		twitchClient: tc,
	}

	go dloader.dlxd(target, link)

	return ""
}

func (dl *downloader) dlxd(target, link string) {
	goutubedl.Path = "yt-dlp"

	dl.twitchClient.Say(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		dl.Log.Errorln(err)
	}
	rExt := "mp4"
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		dl.Log.Errorln(err)
	}
	dl.twitchClient.Say(target, "Downloaded..")
	fn, err := uuid.NewUUID()
	if err != nil {
		dl.Log.Errorln(err)
	}
	f, err := os.Create(fmt.Sprintf("%s.%s", fn, rExt))
	dl.twitchClient.Say(target, fmt.Sprintf("Filename: %s.%s", fn, rExt))

	if err != nil {
		dl.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		dl.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()
	duration := 5 * time.Second
	dl.twitchClient.Say(target, "ResidentSleeper ..")
	time.Sleep(duration)

	dl.upload(target, fmt.Sprintf("%s.%s", fn, rExt))

}

func (dl *downloader) upload(target, path string) {
	const URL = "https://i.yaf.ee/upload"
	dl.twitchClient.Say(target, "Uploading .. dankCircle")
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()

		err := form.WriteField("name", "xd")
		if err != nil {
			dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		file, err := os.Open(path) // path to image file
		if err != nil {
			dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, URL, pr)
	if err != nil {
		dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		dl.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	dl.twitchClient.Say(target, "Uploaded .. PogChamp")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		dl.twitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		dl.Log.Errorln("Error while reading response:", err)
		return
	}

	var reply = string(body[:])

	dl.twitchClient.Say(target, fmt.Sprintf("Removing file: %s", path))
	os.Remove(path)
	dl.twitchClient.Say(target, reply)
}
