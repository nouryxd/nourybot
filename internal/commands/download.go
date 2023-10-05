package commands

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
	"github.com/wader/goutubedl"
	"go.uber.org/zap"
)

type Downloader struct {
	TwitchClient *twitch.Client
	Log          *zap.SugaredLogger
}

func NewDownload(destination, target, link string, tc *twitch.Client, log *zap.SugaredLogger) {
	dl := &Downloader{
		TwitchClient: tc,
		Log:          log,
	}

	switch destination {
	case "catbox":
		dl.CatboxDownload(target, link)
	case "yaf":
		dl.YafDownload(target, link)
	case "kappa":
		dl.KappaDownload(target, link)
	case "gofile":
		dl.GofileDownload(target, link)
	}
}

func (dl *Downloader) YafDownload(target, link string) {
	goutubedl.Path = "yt-dlp"

	dl.TwitchClient.Say(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		dl.Log.Errorln(err)
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		dl.Log.Errorln(err)
	}
	dl.TwitchClient.Say(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		dl.Log.Errorln(err)
	}
	fileName := fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	dl.TwitchClient.Say(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		dl.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		dl.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go NewUpload("yaf", fileName, target, dl.TwitchClient, dl.Log)

}

func (dl *Downloader) KappaDownload(target, link string) {
	goutubedl.Path = "yt-dlp"

	dl.TwitchClient.Say(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		dl.Log.Errorln(err)
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		dl.Log.Errorln(err)
	}
	dl.TwitchClient.Say(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		dl.Log.Errorln(err)
	}
	fileName := fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	dl.TwitchClient.Say(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		dl.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		dl.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go NewUpload("kappa", fileName, target, dl.TwitchClient, dl.Log)

}

func (dl *Downloader) GofileDownload(target, link string) {
	goutubedl.Path = "yt-dlp"

	dl.TwitchClient.Say(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		dl.Log.Errorln(err)
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		dl.Log.Errorln(err)
	}
	dl.TwitchClient.Say(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		dl.Log.Errorln(err)
	}
	fileName := fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	dl.TwitchClient.Say(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		dl.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		dl.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go NewUpload("gofile", fileName, target, dl.TwitchClient, dl.Log)

}

func (dl *Downloader) CatboxDownload(target, link string) {
	goutubedl.Path = "yt-dlp"
	var fileName string

	dl.TwitchClient.Say(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		dl.Log.Errorln(err)
	}

	// I don't know why but I need to set it to mp4, otherwise if
	// I use `result.Into.Ext` catbox won't play the video in the
	// browser and say this message:
	// `No video with supported format and MIME type found.`
	rExt := "mp4"

	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		dl.Log.Errorln(err)
	}
	dl.TwitchClient.Say(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		dl.Log.Errorln(err)
	}
	fileName = fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	dl.TwitchClient.Say(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		dl.Log.Errorln(err)
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		dl.Log.Errorln(err)
	}

	downloadResult.Close()
	f.Close()

	go NewUpload("catbox", fileName, target, dl.TwitchClient, dl.Log)
}
