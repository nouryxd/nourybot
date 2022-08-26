package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/common"
)

// AddTimer slices the message into relevant parts, adding the values onto a
// new data.Timer struct so that the timer can be inserted into the database.
func (app *Application) AddTimer(name string, message twitch.PrivateMessage) {
	cmdParams := strings.SplitN(message.Message, " ", 500)
	// prefixLength is the length of `()addcommand` plus +2 (for the space and zero based)
	prefixLength := 12
	repeat := cmdParams[2]

	// Split the message into the parts we need.
	//
	// message:  ()addtimer   sponsor    20m  hecking love my madmonq pills BatChest
	// parts:    | prefix  |  |name | |repeat | <----------- text ------------->   |
	text := message.Message[prefixLength+len(name)+len(cmdParams[2]) : len(message.Message)]

	// validateTimeFormat will be true if the repeat parameter is in
	// the format of either 30m, 10h, or 10h30m.
	validateTimeFormat, err := regexp.MatchString(`^(\d{1,2}[h])$|^(\d+[m])$|^((\d{1,2}[h])((([0]?|[1-5]{1})[0-9])[m]))$`, repeat)
	if err != nil {
		app.Logger.Errorw("Received malformed time format in timer",
			"repeat", repeat,
			"error", err,
		)
		return
	}

	timer := &data.Timer{
		Name:    name,
		Text:    text,
		Channel: message.Channel,
		Repeat:  repeat,
	}

	// Check if the time string we got is valid, this is important
	// because the Scheduler panics instead of erroring out if an invalid
	// time format string is supplied.
	if validateTimeFormat {
		timer := &data.Timer{
			Name:    name,
			Text:    text,
			Channel: message.Channel,
			Repeat:  repeat,
		}

		err = app.Models.Timers.Insert(timer)
		if err != nil {
			app.Logger.Errorw("Error inserting new timer into database",
				"timer", timer,
				"error", err,
			)

			reply := fmt.Sprintln("Something went wrong FeelsBadMan")
			common.Send(message.Channel, reply, app.TwitchClient)
			return
		} else {
			// cronName is the internal, unique tag/name for the timer.
			// A timer named "sponsor" in channel "forsen" will be named "forsensponsor"
			cronName := fmt.Sprintf("%s%s", message.Channel, name)

			app.Scheduler.AddFunc(fmt.Sprintf("@every %s", repeat), func() { app.newPrivateMessageTimer(message.Channel, text) }, cronName)
			app.Logger.Infow("Added new timer",
				"timer", timer,
			)

			reply := fmt.Sprintf("Successfully added timer %s repeating every %s", name, repeat)
			common.Send(message.Channel, reply, app.TwitchClient)
			return
		}
	} else {
		app.Logger.Errorw("Received malformed time format in timer",
			"timer", timer,
			"error", err,
		)
		reply := fmt.Sprintf("Something went wrong FeelsBadMan received wrong time format. Allowed formats: 30m, 10h, 10h30m")
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// InitialTimers is called on startup and queries the database for a list of
// timers and then adds each onto the scheduler.
func (app *Application) InitialTimers() {
	timer, err := app.Models.Timers.GetAll()
	if err != nil {
		app.Logger.Errorw("Error trying to retrieve all timers from database", err)
		return
	}

	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var ts []*data.Timer

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range timer {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i

		// cronName is the internal, unique tag/name for the timer.
		// A timer named "sponsor" in channel "forsen" will be named "forsensponsor"
		cronName := fmt.Sprintf("%s%s", v.Channel, v.Name)

		// Repeating is at what times the timer should repeat.
		// 2 minute timer is @every 2m
		repeating := fmt.Sprintf("@every %s", v.Repeat)

		// Add new value to the slice
		ts = append(ts, v)

		app.Scheduler.AddFunc(repeating, func() { app.newPrivateMessageTimer(v.Channel, v.Text) }, cronName)
	}

	// Log the initial timers
	app.Logger.Infow("Initial timers",
		"timer",
		ts,
	)
	return
}

// newPrivateMessageTimer is a helper function to set timers
// which trigger into sending a twitch PrivateMessage.
func (app *Application) newPrivateMessageTimer(channel, text string) {
	common.Send(channel, text, app.TwitchClient)
	return
}

// DeleteTimer takes in the name of a timer and tries to delete the timer from the database.
func (app *Application) DeleteTimer(name string, message twitch.PrivateMessage) {
	cronName := fmt.Sprintf("%s%s", message.Channel, name)
	app.Scheduler.RemoveJob(cronName)
	app.Logger.Infow("Deleting timer",
		"name", name,
		"message.Channel", message.Channel,
		"cronName", cronName,
	)

	err := app.Models.Timers.Delete(name)
	if err != nil {
		app.Logger.Errorw("Error deleting timer from database",
			"name", name,
			"cronName", cronName,
			"error", err,
		)

		reply := fmt.Sprintln("Something went wrong FeelsBadMan")
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}

	reply := fmt.Sprintf("Deleted timer with name %s", name)
	common.Send(message.Channel, reply, app.TwitchClient)
	return
}
