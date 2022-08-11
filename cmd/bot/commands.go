package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func (app *Application) AddCommand(name string, message twitch.PrivateMessage) {
	// prefixLength is the length of `()addcommand` plus +2 (for the space and zero based)
	prefixLength := 14

	// Split the twitch message at the length of the prefix + the length of the name of the command.
	//      prefixLength |name| text
	//      0123456789012|4567|
	// e.g. ()addcommand dank FeelsDankMan
	//      |   part1    snip ^  part2   |
	text := message.Message[prefixLength+len(name) : len(message.Message)]
	err := app.Models.Commands.Insert(name, text)

	//	app.Logger.Infow("Message splits",
	//		"Command Name:", name,
	//		"Command Text:", text)

	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	} else {

		reply := fmt.Sprintf("Successfully added command: %s", name)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

func (app *Application) GetCommand(name, username string) (string, error) {
	// Fetch the command from the database if it exists.
	command, err := app.Models.Commands.Get(name)
	if err != nil {
		// It probably did not exist
		return "", err
	}

	// If the command has no level set just return the text.
	// Otherwise check if the level is high enough.
	if command.Level == 0 {
		return command.Text, nil
	} else {
		// Get the user from the database to check if the userlevel is equal
		// or higher than the command.Level.
		user, err := app.Models.Users.Get(username)
		if err != nil {
			return "", err
		}
		if user.Level >= command.Level {
			// Userlevel is sufficient so return the command.Text
			return command.Text, nil
		}
	}

	// Userlevel was not enough so return an empty string and error.
	return "", ErrUserInsufficientLevel

}

func (app *Application) EditCommandLevel(name, lvl string, message twitch.PrivateMessage) {

	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrCommandLevelNotInteger), app.TwitchClient)
		return
	}

	err = app.Models.Commands.SetLevel(name, level)

	if err != nil {
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), app.TwitchClient)
		app.Logger.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated command %s to level %v", name, level)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}
