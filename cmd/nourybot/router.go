package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/lyx0/nourybot/internal/common"
)

func (app *application) startRouter() {
	router := httprouter.New()
	router.GET("/status", app.statusPageRoute)
	router.GET("/commands/:channel", app.channelCommandsRoute)

	app.Log.Fatal(http.ListenAndServe(":8080", router))
}

func (app *application) channelCommandsRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channel := ps.ByName("channel")

	command, err := app.Models.Commands.GetAllChannel(channel)
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return
	}

	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var cs []string

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range command {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var c string

		if v.Category == "ascii" {
			c = fmt.Sprintf(
				"Name: \t%v\n"+
					"Help: \t%v\n"+
					"Level: \t%v\n"+
					"\n",
				v.Name, v.Help, v.Level,
			)
		} else {
			c = fmt.Sprintf(
				"Name: \t%v\n"+
					"Help: \t%v\n"+
					"Level: \t%v\n"+
					"Text: \t%v\n"+
					"\n",
				v.Name, v.Help, v.Level, v.Text,
			)
		}

		// Add new value to the slice
		cs = append(cs, c)

	}

	text := strings.Join(cs, "")
	fmt.Fprintf(w, fmt.Sprint(text), ps.ByName("name"))
}

func (app *application) statusPageRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	commit := common.GetVersion()
	started := common.GetUptime().Format("2006-1-2 15:4:5")
	commitLink := fmt.Sprintf("https://github.com/lyx0/nourybot/commit/%v", common.GetVersionPure())

	fmt.Fprintf(w, fmt.Sprintf("started: \t%v\nenvironment: \t%v\ncommit: \t%v\ngithub: \t%v", started, app.Environment, commit, commitLink))
}
