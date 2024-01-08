package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/wader/goutubedl"
)

func (app *application) NewDownload(destination, target, link string, msg twitch.PrivateMessage) {
	identifier := uuid.NewString()
	go app.Models.Uploads.Insert(
		msg.User.Name,
		msg.User.ID,
		msg.Channel,
		msg.Message,
		destination,
		link,
		identifier,
	)

	switch destination {
	case "catbox":
		app.CatboxDownload(target, link, identifier, msg)
	case "yaf":
		app.YafDownload(target, link, identifier, msg)
	case "kappa":
		app.KappaDownload(target, link, identifier, msg)
	case "gofile":
		app.GofileDownload(target, link, identifier, msg)
	}
}

func (app *application) ConvertAndSave(fName, link string, msg twitch.PrivateMessage) {
	goutubedl.Path = "yt-dlp"
	uuid_og := uuid.NewString()

	app.Send(msg.Channel, "Downloading... dankCircle", msg)
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	//app.Send(target, "Downloaded.", msg)
	fileName := fmt.Sprintf("/public/uploads/%s.%s", uuid_og, rExt)
	f, err := os.Create(fileName)
	//app.Send(target, fmt.Sprintf("Filename: %s", fileName), msg)
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	downloadResult.Close()
	f.Close()

	out := fmt.Sprintf("/public/uploads/%s.mp4", fName)
	fn, err := os.Create(out)
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer fn.Close()

	err = ffmpeg.Input(fileName).
		Output(out).
		OverWriteOutput().ErrorToStdOut().Run()

	defer os.Remove(fileName)
	app.Send(msg.Channel, fmt.Sprintf("https://bot.noury.li/uploads/%s.mp4", fName), msg)
}

func (app *application) ConvertToMP4(link string, msg twitch.PrivateMessage) {
	fName := "in"
	goutubedl.Path = "yt-dlp"
	uuid_og := uuid.NewString()
	uuid_new := uuid.NewString()

	app.Send(msg.Channel, "Downloading... dankCircle", msg)
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	//app.Send(target, "Downloaded.", msg)
	fileName := fmt.Sprintf("/public/uploads/%s.%s", uuid_og, rExt)
	f, err := os.Create(fileName)
	//app.Send(target, fmt.Sprintf("Filename: %s", fileName), msg)
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	downloadResult.Close()
	f.Close()

	out := fmt.Sprintf("/public/uploads/%s.mp4", fName)
	fn, err := os.Create(out)
	if err != nil {
		app.Log.Errorln(err)
		app.Send(msg.Channel, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer fn.Close()

	err = ffmpeg.Input(fileName).
		Output(out).
		OverWriteOutput().ErrorToStdOut().Run()

	defer os.Remove(fileName)
	go app.NewUpload("yaf", out, msg.Channel, uuid_new, msg)
}

func (app *application) YafDownload(target, link, identifier string, msg twitch.PrivateMessage) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle", msg)
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	//app.Send(target, "Downloaded.", msg)
	fileName := fmt.Sprintf("/public/uploads/%s.%s", identifier, rExt)
	f, err := os.Create(fileName)
	//app.Send(target, fmt.Sprintf("Filename: %s", fileName), msg)

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go app.NewUpload("yaf", fileName, target, identifier, msg)

}

func (app *application) KappaDownload(target, link, identifier string, msg twitch.PrivateMessage) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle", msg)
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	//app.Send(target, "Downloaded.", msg)
	fileName := fmt.Sprintf("/public/uploads/%s.%s", identifier, rExt)
	f, err := os.Create(fileName)
	//app.Send(target, fmt.Sprintf("Filename: %s", fileName), msg)

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go app.NewUpload("kappa", fileName, target, identifier, msg)

}

func (app *application) GofileDownload(target, link, identifier string, msg twitch.PrivateMessage) {
	goutubedl.Path = "yt-dlp"

	app.Send(target, "Downloading... dankCircle", msg)
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	safeFilename := fmt.Sprintf("download_%s", result.Info.Title)
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	//app.Send(target, "Downloaded.", msg)
	fileName := fmt.Sprintf("/public/uploads/%s.%s", safeFilename, rExt)
	f, err := os.Create(fileName)
	//app.Send(target, fmt.Sprintf("Filename: %s", fileName), msg)

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	downloadResult.Close()
	f.Close()
	// duration := 5 * time.Second
	// dl.twitchClient.Say(target, "ResidentSleeper ..")
	// time.Sleep(duration)

	go app.NewUpload("gofile", fileName, target, identifier, msg)

}

func (app *application) CatboxDownload(target, link, identifier string, msg twitch.PrivateMessage) {
	goutubedl.Path = "yt-dlp"
	var fileName string

	app.Send(target, "Downloading... dankCircle", msg)
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
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
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	//app.Send(target, "Downloaded.", msg)
	fileName = fmt.Sprintf("/public/uploads/%s.%s", identifier, rExt)
	f, err := os.Create(fileName)
	//app.Send(target, fmt.Sprintf("Filename: %s", fileName), msg)

	if err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, downloadResult); err != nil {
		app.Log.Errorln(err)
		app.Send(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err), msg)
		return
	}

	downloadResult.Close()
	f.Close()

	go app.NewUpload("catbox", fileName, target, identifier, msg)
}
