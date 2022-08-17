package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/common"
)

// AddCommand takes in a name parameter and a twitch.PrivateMessage. It slices the
// twitch.PrivateMessage after the name parameter and adds everything after to a text
// value. Then it calls the app.Models.Commands.Insert method with both name, and text
// values adding them to the database.
func (app *Application) AddCommand(name string, message twitch.PrivateMessage) {
	// prefixLength is the length of `()addcommand` plus +2 (for the space and zero based)
	prefixLength := 14

	// Split the twitch message at the length of the prefix + the length of the name of the command.
	//      prefixLength |name| text
	//      0123456789012|4567|
	// e.g. ()addcommand dank FeelsDankMan
	//      |   part1    snip ^  part2   |
	text := message.Message[prefixLength+len(name) : len(message.Message)]
	command := &data.Command{
		Name:     name,
		Text:     text,
		Category: "uncategorized",
		Level:    0,
	}
	err := app.Models.Commands.Insert(command)

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

// GetCommand queries the database for a name. If an entry exists it checks
// if the Command.Level is 0, if it is the command.Text value is returned.
//
// If the Command.Level is not 0 it queries the database for the level of the
// user who sent the message. If the users level is equal or higher
// the command.Text field is returned.
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

// EditCommandLevel takes in a name and level string and updates the entry with name
// to the supplied level value.
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

// EditCommandCategory takes in a name and category string and updates the command
// in the databse with the passed in new category.
func (app *Application) EditCommandCategory(name, category string, message twitch.PrivateMessage) {
	err := app.Models.Commands.SetCategory(name, category)

	if err != nil {
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), app.TwitchClient)
		app.Logger.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated command %s to category %v", name, category)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// DeleteCommand takes in a name value and deletes the command from the database if it exists.
func (app *Application) DeleteCommand(name string, message twitch.PrivateMessage) {
	err := app.Models.Commands.Delete(name)
	if err != nil {
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
		app.Logger.Error(err)
		return
	}

	reply := fmt.Sprintf("Deleted command %s", name)
	common.Send(message.Channel, reply, app.TwitchClient)
}
