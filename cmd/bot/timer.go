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

		reply := fmt.Sprintf("Successfully added timer %s repeating every %s", name, repeat)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}
