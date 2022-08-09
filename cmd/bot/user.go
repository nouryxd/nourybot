package main

import (
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/commands/decapi"
	"github.com/lyx0/nourybot/pkg/common"
)

func (app *Application) AddUser(login, lvl string, message twitch.PrivateMessage) {
	userId, err := decapi.GetIdByLogin(login)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
		return
	}

	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrUserLevelNotInteger), app.TwitchClient)
		return
	}

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

func (app *Application) EditUserLevel(user, lvl string, message twitch.PrivateMessage) {

	level, err := strconv.Atoi(lvl)
	if err != nil {
		app.Logger.Error(err)
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrUserLevelNotInteger), app.TwitchClient)
		return
	}

	err = app.Models.Users.SetLevel(user, level)

	if err != nil {
		common.Send(message.Channel, fmt.Sprintf("Something went wrong FeelsBadMan %s", ErrRecordNotFound), app.TwitchClient)
		app.Logger.Error(err)
		return
	} else {
		reply := fmt.Sprintf("Updated user %s to level %v", user, level)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}

}
