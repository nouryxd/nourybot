package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/wader/goutubedl"
)

func (app *application) NewDownload(destination, target, link string) {

	switch destination {
	case "catbox":
		app.CatboxDownload(target, link)
	case "yaf":
		app.YafDownload(target, link)
	case "kappa":
		app.KappaDownload(target, link)
	case "gofile":
		app.GofileDownload(target, link)
	}
}

func (app *application) YafDownload(target, link string) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	app.Send(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	fileName := fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	app.Send(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go app.NewUpload("yaf", fileName, target)

}

func (app *application) KappaDownload(target, link string) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	app.Send(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	fileName := fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	app.Send(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go app.NewUpload("kappa", fileName, target)

}

func (app *application) GofileDownload(target, link string) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	safeFilename := fmt.Sprintf("download_%s", result.Info.Title)
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	app.Send(target, "Downloaded.")
	fileName := fmt.Sprintf("%s.%s", safeFilename, rExt)
	f, err := os.Create(fileName)
	app.TwitchClient.Say(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go app.NewUpload("gofile", fileName, target)

}

func (app *application) CatboxDownload(target, link string) {
	goutubedl.Path = "yt-dlp"
	var fileName string

	app.Send(target, "Downloading... dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}

	// I don't know why but I need to set it to mp4, otherwise if
	// I use `result.Into.Ext` catbox won't play the video in the
	// browser and say this message:
	// `No video with supported format and MIME type found.`
	rExt := "mp4"

	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	app.Send(target, "Downloaded.")
	uuidFilename, err := uuid.NewUUID()
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	fileName = fmt.Sprintf("%s.%s", uuidFilename, rExt)
	f, err := os.Create(fileName)
	app.Send(target, fmt.Sprintf("Filename: %s", fileName))

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}

	downloadResult.Close()
	f.Close()

	go app.NewUpload("catbox", fileName, target)
}
