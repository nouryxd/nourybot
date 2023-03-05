package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/commands"
	"github.com/lyx0/nourybot/internal/common"
	"go.uber.org/zap"
)

// AddUser calls GetIdByLogin to get the twitch id of the login name and then adds
// the login name, twitch id and supplied level to the database.
func (app *Application) InitUser(login, twitchId string, message twitch.PrivateMessage) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	_, err := app.Models.Users.Check(twitchId)
	app.Logger.Error(err)
	if err != nil {
		app.Logger.Infow("InitUser: Adding new user:",
			"login: ", login,
			"twitchId: ", twitchId,
		)
		app.Models.Users.Insert(login, twitchId)

		return
	}

	sugar.Infow("User Insert: User already registered: xd",
		"login: ", login,
		"twitchId: ", twitchId,
	)
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
		reply := fmt.Sprintf("User %v: ID %v, Login: %s, TwitchID: %v, Level: %v", login, user.ID, user.Login, user.TwitchID, user.Level)
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

// SetUserLocation sets new location for the user
func (app *Application) SetUserLocation(message twitch.PrivateMessage) {
	// snipLength is the length we need to "snip" off of the start of `message`.
	//  `()set location` = +13
	//  trailing space =  +1
	//      zero-based =  +1
	//                 =  16
	snipLength := 15

	// Split the twitch message at `snipLength` plus length of the name of the
	// The part of the message we are left over with is then passed on to the database
	// handlers as the `location` part of the command.
	location := message.Message[snipLength:len(message.Message)]
	login := message.User.Name
	twitchId := message.User.ID

	app.Logger.Infow("SetUserLocation",
		"location", location,
		"login", login,
		"twitchId", message.User.ID,
	)
	err := app.Models.Users.SetLocation(twitchId, location)
	if err != nil {
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), app.TwitchClient)
		app.Logger.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Successfully set your location to %v", location)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// SetUserLastFM tries to update the database record for the supplied
// login name with the new level.
func (app *Application) SetUserLastFM(lastfmUser string, message twitch.PrivateMessage) {
	login := message.User.Name

	app.Logger.Infow("SetUserLastFM",
		"lastfmUser", lastfmUser,
		"login", login,
	)
	err := app.Models.Users.SetLastFM(login, lastfmUser)
	if err != nil {
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), app.TwitchClient)
		app.Logger.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Successfully set your lastfm username to %v", lastfmUser)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

// GetUserLevel takes in a login name and queries the database for an entry
// with such a name value. If there is one it returns the level value as an integer.
// Returns 0 on an error which is the level for unregistered users.
func (app *Application) GetUserLevel(twitchId string) int {
	userLevel, err := app.Models.Users.GetLevel(twitchId)
	if err != nil {
		return 0
	} else {
		return userLevel
	}
}

func (app *Application) UserCheckWeather(message twitch.PrivateMessage) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	twitchLogin := message.User.Name
	twitchId := message.User.ID
	sugar.Infow("UserCheckWeather: ",
		"twitchLogin:", twitchLogin,
		"twitchId:", twitchId,
	)
	location, err := app.Models.Users.GetLocation(twitchId)
	if err != nil {
		sugar.Errorw("No location data registered for: ",
			"twitchLogin:", twitchLogin,
			"twitchId:", twitchId,
		)
		reply := "No location for your account set in my database. Use ()set location <location> to register. Otherwise use ()weather <location> without registering."
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}

	target := message.Channel
	sugar.Infow("Twitchlogin: ",
		"twitchLogin:", twitchLogin,
		"location:", location,
	)

	commands.Weather(target, location, app.TwitchClient)
}

func (app *Application) UserCheckLastFM(message twitch.PrivateMessage) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	twitchLogin := message.User.Name
	sugar.Infow("Twitchlogin: ",
		"twitchLogin:", twitchLogin,
	)
	lastfmUser, err := app.Models.Users.GetLastFM(twitchLogin)
	if err != nil {
		sugar.Errorw("No LastFM account registered for: ",
			"twitchLogin:", twitchLogin,
		)
		reply := "No lastfm account registered in my database. Use ()register lastfm <username> to register. (Not yet implemented) Otherwise use ()lastfm <username> without registering."
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}

	target := message.Channel
	sugar.Infow("Twitchlogin: ",
		"twitchLogin:", twitchLogin,
		"user:", lastfmUser,
	)

	commands.LastFmUserRecent(target, lastfmUser, app.TwitchClient)
}
