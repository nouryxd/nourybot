package main

import (
	"context"
	"fmt"
	"io"
	"os"

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
	rExt := result.Info.Ext
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Errorln(err)
	}
	app.Send(target, "Downloaded..")
	defer downloadResult.Close()
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

	app.UploadCatbox(target, fmt.Sprintf("%s.%s", fn, rExt))

	//os.Remove(fmt.Sprintf("%s.mp4", fn))

	//app.TwitchClient.Say(channel, b.String())

}
