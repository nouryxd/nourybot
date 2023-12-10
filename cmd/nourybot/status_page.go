package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/pkg/humanize"
)

func (app *application) statusPage() {
	commit := common.GetVersion()
	botUptime := humanize.Time(common.GetUptime())

	statusHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, fmt.Sprintf("up\n\nlast restart:\t %v\nenv: %v\tcommit: \t%v", botUptime, app.Environment, commit))
	}

	http.HandleFunc("/status", statusHandler)
	app.Log.Fatal(http.ListenAndServe(":8080", nil))
}
