package main

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/wader/goutubedl"
)

func (app *application) Download(channel, link string) (string, error) {
	goutubedl.Path = "yt-dlp"

	app.Send(channel, "dankCircle")
	result, err := goutubedl.New(context.Background(), link, goutubedl.Options{})
	if err != nil {
		app.Log.Fatal(err)
	}
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		app.Log.Fatal(err)
	}
	defer downloadResult.Close()
	f, err := os.Create("output.mp4")
	if err != nil {
		app.Log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, downloadResult)

	out, err := exec.Command("curl", "-L", "-F", "file=@output.mp4", "i.yaf.ee/upload").Output()
	if err != nil {
		app.Log.Fatal(err)
	}
	return string(out), nil

}
