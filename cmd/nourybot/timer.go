package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/google/uuid"
	"github.com/lyx0/nourybot/internal/data"
)

// AddTimer slices the message into relevant parts, adding the values onto a
// new data.Timer struct so that the timer can be inserted into the database.
func (app *application) AddTimer(name, repeat string, message twitch.PrivateMessage) {
	cmdParams := strings.SplitN(message.Message, " ", 500)
	// prefixLength is the length of `()add timer` plus +2 (for the space and zero based)
	prefix := "()add timer"
	prefixLength := len("()add timer")
	nameLength := len(name)
	repeatLength := len(repeat)
	app.Log.Infow("Lengths",
		"prefix", prefixLength,
		"name", nameLength,
		"repeat", repeatLength,
		"repeat2", len(cmdParams[2]),
		"sum", prefixLength+nameLength+repeatLength,
	)

	// Split the message into the parts we need.
	//
	// message:  ()timer add sponsor    20m  hecking love my madmonq pills BatChest
	// parts:    | prefix  |  |name | |repeat | <----------- text ------------->   |
	text := message.Message[3+len(prefix)+len(name)+len(repeat) : len(message.Message)]

	// validateTimeFormat will be true if the repeat parameter is in
	// the format of either 30m, 10h, or 10h30m.
	validateTimeFormat, err := regexp.MatchString(`^(\d{1,2}[h])$|^(\d+[m])$|^(\d+[s])$|((\d{1,2}[h])((([0]?|[1-5]{1})[0-9])[m]))$`, repeat)
	if err != nil {
		app.Log.Errorw("Received malformed time format in timer",
			"repeat", repeat,
			"error", err,
		)
		return
	}
	id := uuid.NewString()
	timer := &data.Timer{
		Name:       name,
		Text:       text,
		Identifier: id,
		Channel:    message.Channel,
		Repeat:     repeat,
	}

	// Check if the time string we got is valid, this is important
	// because the Scheduler panics instead of erroring out if an invalid
	// time format string is supplied.
	if validateTimeFormat {
		timer := &data.Timer{
			Name:       name,
			Text:       text,
			Identifier: id,
			Channel:    message.Channel,
			Repeat:     repeat,
		}

		err = app.Models.Timers.Insert(timer)
		if err != nil {
			app.Log.Errorw("Error inserting new timer into database",
				"timer", timer,
				"error", err,
			)

			reply := fmt.Sprintln("Something went wrong FeelsBadMan")
			app.Send(message.Channel, reply)
			return
		} else {
			// cronName is the internal, unique tag/name for the timer.
			// A timer named "sponsor" in channel "forsen" will be named "forsensponsor"

			app.Scheduler.AddFunc(fmt.Sprintf("@every %s", repeat), func() { app.newPrivateMessageTimer(message.Channel, text) }, id)
			app.Log.Infow("Added new timer",
				"timer", timer,
			)

			reply := fmt.Sprintf("Successfully added timer %s repeating every %s", name, repeat)
			app.Send(message.Channel, reply)
			return
		}
	} else {
		app.Log.Errorw("Received malformed time format in timer",
			"timer", timer,
			"error", err,
		)
		reply := "Something went wrong FeelsBadMan received wrong time format. Allowed formats: 30m, 10h, 10h30m"
		app.Send(message.Channel, reply)
		return
	}
}

// EditTimer just contains the logic for deleting a timer, and then adding a new one
// with the same name. It is technically not editing the timer.
func (app *application) EditTimer(name, repeat string, message twitch.PrivateMessage) {
	// Check if a timer with that name is in the database.
	app.Log.Info(name)

	old, err := app.Models.Timers.Get(name)
	if err != nil {
		app.Log.Errorw("Could not get timer",
			"timer", old,
			"error", err,
		)
		reply := "Something went wrong FeelsBadMan"
		app.Send(message.Channel, reply)
		return
	}

	// -----------------------
	// Delete the old timer
	// -----------------------
	identifier := old.Identifier
	app.Scheduler.RemoveJob(identifier)

	err = app.Models.Timers.Delete(name)
	if err != nil {
		app.Log.Errorw("Error deleting timer from database",
			"name", name,
			"identifier", identifier,
			"error", err,
		)

		reply := fmt.Sprintln("Something went wrong FeelsBadMan")
		app.Send(message.Channel, reply)
		return
	}

	// -----------------------
	// Add the new timer
	// -----------------------
	//cmdParams := strings.SplitN(message.Message, " ", 500)
	// prefixLength is the length of `()editcommand` plus +2 (for the space and zero based)
	prefix := "()edit timer"
	prefixLength := len("()add timer")
	nameLength := len(name)
	repeatLength := len(repeat)
	app.Log.Infow("Lengths",
		"prefix", prefixLength,
		"name", nameLength,
		"repeat", repeatLength,
		"sum", prefixLength+nameLength+repeatLength,
	)

	// Split the message into the parts we need.
	//
	// message:  ()addtimer   sponsor    20m  hecking love my madmonq pills BatChest
	// parts:    | prefix  |  |name | |repeat | <----------- text ------------->   |
	text := message.Message[3+len(prefix)+len(name)+len(repeat) : len(message.Message)]

	// validateTimeFormat will be true if the repeat parameter is in
	// the format of either 30m, 10h, or 10h30m.
	validateTimeFormat, err := regexp.MatchString(`^(\d{1,2}[h])$|^(\d+[m])$|^(\d+[s])$|((\d{1,2}[h])((([0]?|[1-5]{1})[0-9])[m]))$`, repeat)
	if err != nil {
		app.Log.Errorw("Received malformed time format in timer",
			"repeat", repeat,
			"error", err,
		)
		return
	}

	id := uuid.NewString()
	timer := &data.Timer{
		Name:       name,
		Text:       text,
		Identifier: id,
		Channel:    message.Channel,
		Repeat:     repeat,
	}

	// Check if the time string we got is valid, this is important
	// because the Scheduler panics instead of erroring out if an invalid
	// time format string is supplied.
	if validateTimeFormat {
		timer := &data.Timer{
			Name:       name,
			Text:       text,
			Identifier: id,
			Channel:    message.Channel,
			Repeat:     repeat,
		}

		err = app.Models.Timers.Insert(timer)
		if err != nil {
			app.Log.Errorw("Error inserting new timer into database",
				"timer", timer,
				"error", err,
			)

			reply := fmt.Sprintln("Something went wrong FeelsBadMan")
			app.Send(message.Channel, reply)
			return
		} else { // this is a bit scuffed. The else here is the end of a successful call.
			// cronName is the internal, unique tag/name for the timer.
			// A timer named "sponsor" in channel "forsen" will be named "forsensponsor"
			app.Scheduler.AddFunc(fmt.Sprintf("@every %s", repeat), func() { app.newPrivateMessageTimer(message.Channel, text) }, id)

			app.Log.Infow("Updated a timer",
				"Name", name,
				"Channel", message.Channel,
				"Old timer", old,
				"New timer", timer,
			)

			reply := fmt.Sprintf("Successfully updated timer %s", name)
			app.Send(message.Channel, reply)
			return
		}
	} else {
		app.Log.Errorw("Received malformed time format in timer",
			"timer", timer,
			"error", err,
		)
		reply := "Something went wrong FeelsBadMan received wrong time format. Allowed formats: 30s, 30m, 10h, 10h30m"
		app.Send(message.Channel, reply)
		return
	}
}

// InitialTimers is called on startup and queries the database for a list of
// timers and then adds each onto the scheduler.
func (app *application) ListTimers() string {
	timer, err := app.Models.Timers.GetAll()
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}

	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var ts []string

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range timer {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i

		// cronName is the internal, unique tag/name for the timer.
		// A timer named "sponsor" in channel "forsen" will be named "forsensponsor"
		t := fmt.Sprintf(
			"ID: \t\t%v\n"+
				"Name: \t\t%v\n"+
				"Identifier: \t%v\n"+
				"Text: \t\t%v\n"+
				"Channel: \t%v\n"+
				"Repeat: \t%v\n"+
				"\n\n",
			v.ID, v.Name, v.Identifier, v.Text, v.Channel, v.Repeat,
		)

		// Add new value to the slice
		ts = append(ts, t)

		//app.Scheduler.AddFunc(repeating, func() { app.newPrivateMessageTimer(v.Channel, v.Text) }, cronName)
	}

	reply, err := app.uploadPaste(strings.Join(ts, ""))
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
		return ""
	}

	return reply
}

// InitialTimers is called on startup and queries the database for a list of
// timers and then adds each onto the scheduler.
func (app *application) InitialTimers() {
	timer, err := app.Models.Timers.GetAll()
	if err != nil {
		app.Log.Errorw("Error trying to retrieve all timers from database", err)
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

		// Repeating is at what times the timer should repeat.
		// 2 minute timer is @every 2m
		repeating := fmt.Sprintf("@every %s", v.Repeat)

		// Add new value to the slice
		ts = append(ts, v)

		app.Scheduler.AddFunc(repeating, func() { app.newPrivateMessageTimer(v.Channel, v.Text) }, v.Identifier)
	}

	app.Log.Infow("Initial timers",
		"timer", ts,
	)
}

// newPrivateMessageTimer is a helper function to set timers
// which trigger into sending a twitch PrivateMessage.
func (app *application) newPrivateMessageTimer(channel, text string) {
	app.Send(channel, text)
}

// DeleteTimer takes in the name of a timer and tries to delete the timer from the database.
func (app *application) DeleteTimer(name string, message twitch.PrivateMessage) {

	identifier, err := app.Models.Timers.GetIdentifier(name)
	if err != nil {
		app.Log.Errorw("Error retrieving identifier rom database",
			"name", name,
			"identifier", identifier,
			"error", err,
		)
	}

	app.Scheduler.RemoveJob(identifier)

	app.Log.Infow("Deleting timer",
		"name", name,
		"identifier", identifier,
		"message.Channel", message.Channel,
	)

	err = app.Models.Timers.Delete(identifier)
	if err != nil {
		app.Log.Errorw("Error deleting timer from database",
			"name", name,
			"identifier", identifier,
			"error", err,
		)

		reply := fmt.Sprintln("Something went wrong FeelsBadMan")
		app.Send(message.Channel, reply)
		return
	}

	reply := fmt.Sprintf("Deleted timer with name %s", name)
	app.Send(message.Channel, reply)
}
