package main

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/common"
)

// AddCommand takes in a name parameter and a twitch.PrivateMessage. It slices the
// twitch.PrivateMessage after the name parameter and adds everything after to a text
// value. Then it calls the app.Models.Commands.Insert method with both name, and text
// values adding them to the database.
func (app *Application) AddTimer(name string, message twitch.PrivateMessage) {
	// prefixLength is the length of `()addtimer` plus +2 (for the space and zero based)
	cmdParams := strings.SplitN(message.Message, " ", 500)
	prefixLength := 12
	repeat := cmdParams[2]

	// Split the twitch message at the length of the prefix + the length of the name of the command.
	//      prefixLength |name| text
	//      0123456789012|4567|
	// e.g. ()addcommand dank FeelsDankMan
	//      |   part1    snip ^  part2   |
	text := message.Message[prefixLength+len(name)+len(cmdParams[2]) : len(message.Message)]

	// ()addtimer gfuel 5m Yo buy my cool gfuel cause its cool n shit
	// |         | name |
	timer := &data.Timer{
		Name:    name,
		Text:    text,
		Channel: message.Channel,
		Repeat:  repeat,
	}

	app.Logger.Infow("timer", timer)
	err := app.Models.Timers.Insert(timer)
	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	} else {
		//                         app.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// app.Scheduler.AddFunc(fmt.Sprintf("@every %s", repeat), (func() {
		// app.Scheduler.AddFunc(fmt.Sprintf("@every %s", repeat), func(message.Channel, text) { app.newTimer(message.Channel, text) }))
		cronName := fmt.Sprintf("%s-%s", message.Channel, name)
		app.Scheduler.AddFunc(fmt.Sprintf("@every %s", repeat), func() { app.newTimer(message.Channel, text) }, cronName)

		//app.Scheduler.Tag(fmt.Sprintf("%s-%s", message.Channel, name)).Every(repeat).StartAt(time.Now()).Do(app.newTimer, message.Channel, text
		reply := fmt.Sprintf("Successfully added timer %s repeating every %s (ID: xd)", name, repeat)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// InitialJoin is called on startup and queries the database for a list of
// channels which the TwitchClient then joins.
func (app *Application) InitialTimers() {
	// GetJoinable returns a slice of channel names.
	timer, err := app.Models.Timers.GetAll()
	if err != nil {
		app.Logger.Error(err)
		return
	}

	app.Logger.Info(timer)

	// Iterate over each timer and add them to the scheduler.
	for _, v := range timer {
		cronName := fmt.Sprintf("%s-%s", v.Channel, v.Name)
		app.Scheduler.AddFunc(fmt.Sprintf("@every %s", v.Repeat), func() { app.newTimer(v.Channel, v.Text) }, cronName)
		app.Logger.Infow("Initial timers:",
			"Name", v.Name,
			"Channel", v.Channel,
			"Text", v.Text,
			"Repeat", v.Repeat,
			"V", v,
			"cronName", cronName,
		)

		// app.Scheduler.Tag(fmt.Sprintf("%s-%s", v.Channel, v.Name)).Every(v.Repeat).StartAt(time.Now()).Do(app.newTimer, v.Channel, v.Text)

	}
	for _, v := range app.Scheduler.Entries() {
		app.Logger.Info(v)
	}

	//ent := app.Scheduler.Enries()
	app.Logger.Infow("Entries", app.Scheduler.Entries())
}

func (app *Application) newTimer(channel, text string) {
	common.Send(channel, text, app.TwitchClient)
}

// DeleteCommand takes in a name value and deletes the command from the database if it exists.
func (app *Application) DeleteTimer(name string, message twitch.PrivateMessage) {

	cronName := fmt.Sprintf("%s-%s", message.Channel, name)
	app.Scheduler.RemoveJob(cronName)

	// app.Scheduler.Remove(timer.ID)
	err := app.Models.Timers.Delete(name)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
	}

	reply := fmt.Sprintf("Deleted timer %s", name)
	common.Send(message.Channel, reply, app.TwitchClient)
}
