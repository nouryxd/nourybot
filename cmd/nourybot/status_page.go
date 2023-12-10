package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lyx0/nourybot/internal/common"
)

func (app *application) statusPage() {
	commit := common.GetVersion()
	started := common.GetUptime().Format("2006-1-2 15:4:5")
	commitLink := fmt.Sprintf("https://github.com/lyx0/nourybot/commit/%v", common.GetVersionPure())

	statusHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, fmt.Sprintf("started: \t%v\nenvironment: \t%v\ncommit: \t%v\ngithub: \t%v", started, app.Environment, commit, commitLink))
	}

	http.HandleFunc("/status", statusHandler)
	app.Log.Fatal(http.ListenAndServe(":8080", nil))
}
