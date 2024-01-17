package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/internal/data"
)

func (app *application) startRouter() {
	router := httprouter.New()
	router.GET("/", app.homeRoute)
	router.GET("/status", app.statusPageRoute)
	router.GET("/commands/:channel", app.channelCommandsRoute)
	router.GET("/commands", app.commandsRoute)
	router.GET("/timer/:channel", app.channelTimersRoute)

	// Serve files uploaded by the meme command, but don't list the directory contents.
	fs := justFilesFilesystem{http.Dir("/public/uploads/")}
	router.Handler("GET", "/uploads/*filepath", http.StripPrefix("/uploads", http.FileServer(fs)))

	app.Log.Fatal(http.ListenAndServe(":8080", router))
}

type commandsRouteData struct {
	Commands map[string]command
}

func (app *application) commandsRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("./web/templates/commands.page.gohtml")
	if err != nil {
		app.Log.Error(err)
		return
	}

	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var cs []string

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range helpText {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var c string

		if v.Alias == nil {
			c = fmt.Sprintf("Name: %s\nDescription: %s\nLevel: %s\nUsage: %s\n\n", i, v.Description, v.Level, v.Usage)
		} else {
			c = fmt.Sprintf("Name: %s\nAliases: %s\nDescription: %s\nLevel: %s\nUsage: %s\n\n", i, v.Alias, v.Description, v.Level, v.Usage)

		}

		// Add new value to the slice
		cs = append(cs, c)
	}

	sort.Strings(cs)
	data := &commandsRouteData{helpText}

	err = t.Execute(w, data)
	if err != nil {
		app.Log.Error(err)
		return
	}

}

type homeRouteData struct {
	Channels []*data.Channel
}

func (app *application) homeRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("./web/templates/home.page.gohtml")
	if err != nil {
		app.Log.Error(err)
		return
	}

	allChannel, err := app.Models.Channels.GetAll()
	if err != nil {
		app.Log.Error(err)
		return
	}
	app.Log.Infow("All channels:",
		"channel", allChannel)
	data := &homeRouteData{allChannel}

	err = t.Execute(w, data)
	if err != nil {
		app.Log.Error(err)
		return
	}

}

func (app *application) channelCommandsRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channel := ps.ByName("channel")
	var cs []string
	var text string

	command, err := app.Models.Commands.GetAllChannel(channel)
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all commands for a channel from database", err)
		return
	}
	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.

	heading := fmt.Sprintf("Commands in %s\n\n", channel)
	cs = append(cs, heading)
	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range command {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var c string

		c = fmt.Sprintf(
			"Name: %v\n"+
				"Description: %v\n"+
				"Level: %v\n"+
				"Text: %v\n"+
				"\n",
			v.Name, v.Description, v.Level, v.Text,
		)

		// Add new value to the slice
		cs = append(cs, c)
	}

	text = strings.Join(cs, "")
	fmt.Fprintf(w, fmt.Sprint(text))

}

func (app *application) channelTimersRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channel := ps.ByName("channel")
	var ts []string
	var text string

	timer, err := app.Models.Timers.GetChannelTimer(channel)
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timer for a channel from database", err)
		return
	}
	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range timer {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var t string

		t = fmt.Sprintf(
			"Name: \t%v\n"+
				"Text: \t%v\n"+
				"Repeat: %v\n"+
				"\n",
			v.Name, v.Text, v.Repeat,
		)

		// Add new value to the slice
		ts = append(ts, t)
	}

	text = strings.Join(ts, "")
	fmt.Fprintf(w, fmt.Sprint(text))
}

func (app *application) statusPageRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	commit := common.GetVersion()
	started := common.GetUptime().Format("2006-1-2 15:4:5")
	commitLink := fmt.Sprintf("https://github.com/lyx0/nourybot/commit/%v", common.GetVersionPure())

	fmt.Fprintf(w, fmt.Sprintf("started: \t%v\nenvironment: \t%v\ncommit: \t%v\ngithub: \t%v", started, app.Environment, commit, commitLink))
}

// Since I want to serve the files that I uploaded with the meme command to the /public/uploads
// folder, but not list the directory on the `/uploads/` route I found this issue that solves
// that problem with the httprouter.
//
//	https://github.com/julienschmidt/httprouter/issues/25#issuecomment-74977940
type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
