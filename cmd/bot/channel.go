package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/commands/decapi"
	"github.com/lyx0/nourybot/pkg/common"
)

func AddChannel(login string, message twitch.PrivateMessage, app *Application) {
	userId, err := decapi.GetIdByLogin(login)
	if err != nil {
		app.Logger.Error(err)
		return
	}

	channel := &data.Channel{
		Login:    login,
		TwitchID: userId,
		Announce: false,
	}

	err = app.Models.Channels.Insert(channel)
	if err != nil {
		app.Logger.Error(err)
	}

	reply := fmt.Sprintf("Joined channel %s", login)
	common.Send(message.Channel, reply, app.TwitchClient)
}

func DeleteChannel(login string, message twitch.PrivateMessage, app *Application) {
	err := app.Models.Channels.Delete(login)
	if err != nil {
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
		app.Logger.Error(err)
		return
	}

	reply := fmt.Sprintf("Deleted channel %s", login)
	common.Send(message.Channel, reply, app.TwitchClient)
}
