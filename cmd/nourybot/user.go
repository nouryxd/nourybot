package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/nouryxd/nourybot/pkg/lastfm"
	"github.com/nouryxd/nourybot/pkg/owm"
)

// InitUser is called on each command usage and checks if the user that sent the command
// is already in the database. If he isn't then add the user to the database.
func (app *application) InitUser(login, twitchId string) {
	_, err := app.Models.Users.Check(twitchId)
	//app.Log.Error(err)
	if err != nil {
		go app.Models.Users.Insert(login, twitchId)

		return
	}
}

// DebugUser queries the database for a login name, if that name exists it creates a new paste
func (app *application) DebugUser(login string, message twitch.PrivateMessage) {
	user, err := app.Models.Users.Get(login)

	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		app.Send(message.Channel, reply, message)
		return
	} else {
		// subject := fmt.Sprintf("DEBUG for user %v", login)
		body := fmt.Sprintf("id=%v \nlogin=%v \nlevel=%v \nlocation=%v \nlastfm=%v",
			user.TwitchID,
			user.Login,
			user.Level,
			user.Location,
			user.LastFMUsername,
		)

		resp, err := app.uploadPaste(body)
		if err != nil {
			app.Log.Errorln("Could not upload paste:", err)
			app.Send(message.Channel,
				fmt.Sprintf("Something went wrong FeelsBadMan %v", ErrDuringPasteUpload), message)
			return
		}
		app.Send(message.Channel, resp, message)
		return
	}
}

// DeleteUser takes in a login string, queries the database for an entry with
// that login name and tries to delete that entry in the database.
func (app *application) DeleteUser(login string, message twitch.PrivateMessage) {
	err := app.Models.Users.Delete(login)
	if err != nil {
		app.Send(message.Channel, "Something went wrong FeelsBadMan", message)
		app.Log.Error(err)
		return
	}

	reply := fmt.Sprintf("Deleted user %s", login)
	app.Send(message.Channel, reply, message)
}

// EditUserLevel tries to update the database record for the supplied
// login name with the new level.
func (app *application) EditUserLevel(login, lvl string, message twitch.PrivateMessage) {
	// Convert the level string to an integer. This is an easy check to see if
	// the level supplied was a number only.
	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Log.Error(err)
		app.Send(message.Channel,
			fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrUserLevelNotInteger), message)
		return
	}

	err = app.Models.Users.SetLevel(login, level)
	if err != nil {
		app.Send(message.Channel,
			fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), message)
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated user %s to level %v", login, level)
		app.Send(message.Channel, reply, message)
		return
	}
}

// SetUserLocation sets new location for the user
func (app *application) SetUserLocation(message twitch.PrivateMessage) {
	// snipLength is the length we need to "snip" off of the start of `message`.
	snipLength := 15

	// Split the twitch message at `snipLength` plus length of the name of the
	// The part of the message we are left over with is then passed on to the database
	// handlers as the `location` part of the command.
	location := message.Message[snipLength:len(message.Message)]
	twitchId := message.User.ID

	err := app.Models.Users.SetLocation(twitchId, location)
	if err != nil {
		app.Send(message.Channel,
			fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), message)

		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Successfully set your location to %v", location)
		app.Send(message.Channel, reply, message)
		return
	}
}

// SetUserLastFM tries to update the database record for the supplied
// login name with the new level.
func (app *application) SetUserLastFM(lastfmUser string, message twitch.PrivateMessage) {
	login := message.User.Name

	err := app.Models.Users.SetLastFM(login, lastfmUser)
	if err != nil {
		app.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), message)
		app.Log.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Successfully set your lastfm username to %v", lastfmUser)
		app.Send(message.Channel, reply, message)
		return
	}
}

// GetUserLevel returns the level a user has in the database.
func (app *application) GetUserLevel(msg twitch.PrivateMessage) int {
	var dbUserLevel int
	var twitchUserLevel int

	lvl, err := app.Models.Users.GetLevel(msg.User.ID)
	if err != nil {
		dbUserLevel = 0
	} else {
		dbUserLevel = lvl
	}
	if msg.User.Badges["moderator"] == 1 ||
		msg.User.Badges["broadcaster"] == 1 {
		twitchUserLevel = 250
	} else if msg.User.Badges["vip"] == 1 {
		twitchUserLevel = 100
	} else {
		twitchUserLevel = 0
	}

	if dbUserLevel > twitchUserLevel {
		return dbUserLevel
	} else {
		return twitchUserLevel
	}
}

// UserCheckWeather checks if a user is in the database and if he has a location
// provided. If both is true it calls owm.Weather with the location and replies
// with the result.
// If no location was provided the response will instruct the user
// how to set a location.
func (app *application) UserCheckWeather(message twitch.PrivateMessage) {
	target := message.Channel
	twitchLogin := message.User.Name
	twitchId := message.User.ID

	location, err := app.Models.Users.GetLocation(twitchId)
	if err != nil {
		app.Log.Errorw("No location data registered for: ",
			"twitchLogin:", twitchLogin,
			"twitchId:", twitchId,
		)
		reply := "No location for your account set in my database. Use ()set location <location> to register. Otherwise use ()weather <location> without registering."
		app.Send(message.Channel, reply, message)
		return
	}

	reply, _ := owm.Weather(location)
	app.Send(target, reply, message)
}

// UserCheckWeather checks if a user is in the database and if he has a lastfm
// username provided. If both is true it calls lastfm.LastFmUserRecent with the username
// and replies with the result.
// If no lastfm username was provided the response will instruct the user
// how to set a lastfm username.
func (app *application) UserCheckLastFM(message twitch.PrivateMessage) string {
	twitchLogin := message.User.Name
	target := message.Channel

	lastfmUser, err := app.Models.Users.GetLastFM(twitchLogin)
	if err != nil {
		app.Log.Errorw("No LastFM account registered for: ",
			"twitchLogin:", twitchLogin,
		)
		reply := "No lastfm account registered in my database. Use ()set lastfm <username> to register. Otherwise use ()lastfm <username> without registering."
		return reply
	}

	reply := lastfm.LastFmUserRecent(target, lastfmUser)
	return reply
}
