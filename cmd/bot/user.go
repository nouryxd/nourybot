package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/commands/decapi"
	"github.com/lyx0/nourybot/pkg/common"
)

// AddUser calls GetIdByLogin to get the twitch id of the login name and then adds
// the login name, twitch id and supplied level to the database.
func (app *Application) AddUser(login, lvl string, message twitch.PrivateMessage) {
	userId, err := decapi.GetIdByLogin(login)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
		return
	}

	// Convert the level string to an integer. This is an easy check to see if
	// the level supplied was a number only.
	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrUserLevelNotInteger), app.TwitchClient)
		return
	}

	// Create a user to hold our values to be inserted to the database.
	user := &data.User{
		Login:    login,
		TwitchID: userId,
		Level:    level,
	}

	err = app.Models.Users.Insert(user)
	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	} else {
		reply := fmt.Sprintf("Added user %s with level %v", login, level)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// DebugUser queries the database for a login name, if that name exists it returns the fields
// and outputs them to twitch chat and a twitch whisper.
func (app *Application) DebugUser(login string, message twitch.PrivateMessage) {
	user, err := app.Models.Users.Get(login)

	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	} else {
		reply := fmt.Sprintf("User %v: ID %v, Login: %s, TwitchID: %v, Level %v", login, user.ID, user.Login, user.TwitchID, user.Level)
		common.Send(message.Channel, reply, app.TwitchClient)
		app.TwitchClient.Whisper(message.User.Name, reply)
		return
	}
}

// DeleteUser takes in a login string, queries the database for an entry with
// that login name and tries to delete that entry in the database.
func (app *Application) DeleteUser(login string, message twitch.PrivateMessage) {
	err := app.Models.Users.Delete(login)
	if err != nil {
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
		app.Logger.Error(err)
		return
	}

	reply := fmt.Sprintf("Deleted user %s", login)
	common.Send(message.Channel, reply, app.TwitchClient)
}

// EditUserLevel tries to update the database record for the supplied
// login name with the new level.
func (app *Application) EditUserLevel(login, lvl string, message twitch.PrivateMessage) {
	// Convert the level string to an integer. This is an easy check to see if
	// the level supplied was a number only.
	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrUserLevelNotInteger), app.TwitchClient)
		return
	}

	err = app.Models.Users.SetLevel(login, level)
	if err != nil {
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), app.TwitchClient)
		app.Logger.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated user %s to level %v", login, level)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// GetUserLevel takes in a login name and queries the database for an entry
// with such a name value. If there is one it returns the level value as an integer.
// Returns 0 on an error which is the level for unregistered users.
func (app *Application) GetUserLevel(login string) int {
	user, err := app.Models.Users.Get(login)
	if err != nil {
		return 0
	} else {
		return user.Level
	}
}
