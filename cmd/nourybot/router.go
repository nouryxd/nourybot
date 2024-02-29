package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/julienschmidt/httprouter"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/common"
	"github.com/lyx0/nourybot/pkg/ivr"
	"github.com/nicklaw5/helix/v2"
)

func (app *application) startRouter() {
	router := httprouter.New()
	router.GET("/", app.homeRoute)
	router.GET("/status", app.statusPageRoute)
	router.POST("/eventsub/:channel", app.eventsubFollow)
	router.GET("/commands", app.commandsRoute)
	router.GET("/commands/:channel", app.channelCommandsRoute)
	router.GET("/timer", app.timersRoute)
	router.GET("/timer/:channel", app.channelTimersRoute)

	// Serve files uploaded by the meme command, but don't list the directory contents.
	fs := justFilesFilesystem{http.Dir("/public/uploads/")}
	router.Handler("GET", "/uploads/*filepath", http.StripPrefix("/uploads", http.FileServer(fs)))

	app.Log.Info("Serving on :8080")
	app.Log.Fatal(http.ListenAndServe(":8080", router))
}

type eventSubNotification struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Challenge    string                     `json:"challenge"`
	Event        json.RawMessage            `json:"event"`
}

// eventsubMessageId stores the last message id of an event sub. Twitch resends events
// if it is unsure that you have gotten them so we check if the last event has the same
// message id and if it does discard the event.
var lastEventSubSubscriptionId = []string{"xd"}

func (app *application) eventsubFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channel := ps.ByName("channel")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	// verify that the notification came from twitch using the secret.
	if !helix.VerifyEventSubNotification(app.Config.eventSubSecret, r.Header, string(body)) {
		log.Println("no valid signature on subscription")
		return
	} else {
		log.Println("verified signature for subscription")
	}

	var vals eventSubNotification
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&vals)
	if err != nil {
		log.Println(err)
		return
	}
	// if there's a challenge in the request, respond with only the challenge to verify your eventsub.
	if vals.Challenge != "" {
		w.Write([]byte(vals.Challenge))
		return
	}
	//r.Body.Close()

	// Check if the current events subscription id equals the last events.
	// If it does ignore the event since it's a repeated event.
	for i := 0; i < len(lastEventSubSubscriptionId); i++ {
		if vals.Subscription.ID == lastEventSubSubscriptionId[i] {
			return
		} else {
			lastEventSubSubscriptionId[i] = vals.Subscription.ID
		}
	}

	switch vals.Subscription.Type {
	case helix.EventSubTypeStreamOnline:
		var liveEvent helix.EventSubStreamOnlineEvent

		err = json.NewDecoder(bytes.NewReader(vals.Event)).Decode(&liveEvent)
		log.Printf("got stream online event webhook: [%s]: %s is live\n", channel, liveEvent.BroadcasterUserName)
		w.WriteHeader(200)
		w.Write([]byte("ok"))

		game := ivr.GameByUsername(liveEvent.BroadcasterUserLogin)
		title := ivr.TitleByUsername(liveEvent.BroadcasterUserLogin)

		app.SendNoBanphrase(channel, fmt.Sprintf("%s went live FeelsGoodMan Game: %s; Title: %s; https://twitch.tv/%s", liveEvent.BroadcasterUserName, game, title, liveEvent.BroadcasterUserLogin))

	case helix.EventSubTypeStreamOffline:
		var offlineEvent helix.EventSubStreamOfflineEvent
		err = json.NewDecoder(bytes.NewReader(vals.Event)).Decode(&offlineEvent)
		log.Printf("got stream online event webhook: [%s]: %s is live\n", channel, offlineEvent.BroadcasterUserName)
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		app.SendNoBanphrase(channel, fmt.Sprintf("%s went offline FeelsBadMan", offlineEvent.BroadcasterUserName))
	}
}

type timersRouteData struct {
	Timers []data.Timer
}

func (app *application) timersRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		"./web/templates/base.template.gohtml",
		"./web/templates/header.partial.gohtml",
		"./web/templates/footer.partial.gohtml",
		"./web/templates/timers.page.gohtml",
	)
	if err != nil {
		app.Log.Error(err)
		return
	}
	var ts []data.Timer

	timerData, err := app.Models.Timers.GetAll()
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timer for a channel from database", err)
		return
	}

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range timerData {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var t data.Timer
		t.Name = v.Name
		t.Text = v.Text
		t.Repeat = v.Repeat

		// Add new value to the slice
		ts = append(ts, t)
	}

	data := &timersRouteData{ts}
	err = t.Execute(w, data)
	if err != nil {
		app.Log.Error(err)
		return
	}
}

type channelTimersRouteData struct {
	Timers  []data.Timer
	Channel string
}

func (app *application) channelTimersRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channel := ps.ByName("channel")
	t, err := template.ParseFiles(
		"./web/templates/base.template.gohtml",
		"./web/templates/header.partial.gohtml",
		"./web/templates/footer.partial.gohtml",
		"./web/templates/channeltimers.page.gohtml",
	)
	if err != nil {
		app.Log.Error(err)
		return
	}
	var ts []data.Timer

	timerData, err := app.Models.Timers.GetChannelTimer(channel)
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timer for a channel from database", err)
		return
	}

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range timerData {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var t data.Timer
		t.Name = v.Name
		t.Text = v.Text
		t.Repeat = v.Repeat

		// Add new value to the slice
		ts = append(ts, t)
	}

	data := &channelTimersRouteData{ts, channel}
	err = t.Execute(w, data)
	if err != nil {
		app.Log.Error(err)
		return
	}
}

type commandsRouteData struct {
	Commands map[string]command
}

func (app *application) commandsRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		"./web/templates/base.template.gohtml",
		"./web/templates/header.partial.gohtml",
		"./web/templates/footer.partial.gohtml",
		"./web/templates/commands.page.gohtml",
	)
	if err != nil {
		app.Log.Error(err)
		return
	}

	var cs []string

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
	t, err := template.ParseFiles(
		"./web/templates/base.template.gohtml",
		"./web/templates/header.partial.gohtml",
		"./web/templates/footer.partial.gohtml",
		"./web/templates/home.page.gohtml",
	)

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

type channelCommandsRouteData struct {
	Commands []data.Command
	Channel  string
}

func (app *application) channelCommandsRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	channel := ps.ByName("channel")
	t, err := template.ParseFiles(
		"./web/templates/base.template.gohtml",
		"./web/templates/header.partial.gohtml",
		"./web/templates/footer.partial.gohtml",
		"./web/templates/channelcommands.page.gohtml",
	)
	if err != nil {
		app.Log.Error(err)
		return
	}
	var cs []data.Command

	commandData, err := app.Models.Commands.GetAllChannel(channel)
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timer for a channel from database", err)
		return
	}

	for i, v := range commandData {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var c data.Command
		c.Name = v.Name
		c.Level = v.Level
		c.Description = v.Description
		c.Text = v.Text

		cs = append(cs, c)
	}

	data := &channelCommandsRouteData{cs, channel}
	err = t.Execute(w, data)
	if err != nil {
		app.Log.Error(err)
		return
	}
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
