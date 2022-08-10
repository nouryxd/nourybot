package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/commands/decapi"
	"github.com/lyx0/nourybot/pkg/common"
)

func (app *Application) AddChannel(login string, message twitch.PrivateMessage) {
	userId, err := decapi.GetIdByLogin(login)
	if err != nil {
		app.Logger.Error(err)
		return
	}

	channel := &data.Channel{
		Login:    login,
		TwitchID: userId,
	}

	err = app.Models.Channels.Insert(channel)
	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	} else {
		app.TwitchClient.Join(login)
		reply := fmt.Sprintf("Added channel %s", login)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

func (app *Application) GetAllChannels() {
	channel, err := app.Models.Channels.GetAll()
	if err != nil {
		app.Logger.Error(err)
		return
	}
	app.Logger.Infow("All channels:",
		"channel", channel)
}

func (app *Application) DeleteChannel(login string, message twitch.PrivateMessage) {
	err := app.Models.Channels.Delete(login)
	if err != nil {
		common.Send(message.Channel, "Something went wrong FeelsBadMan", app.TwitchClient)
		app.Logger.Error(err)
		return
	}

	app.TwitchClient.Depart(login)

	reply := fmt.Sprintf("Deleted channel %s", login)
	common.Send(message.Channel, reply, app.TwitchClient)
}

func (app *Application) InitialJoin() {
	channel, err := app.Models.Channels.GetJoinable()
	if err != nil {
		app.Logger.Error(err)
		return
	}

	for _, v := range channel {
		app.TwitchClient.Join(v)
		app.Logger.Infow("Joining channel:",
			"channel", v)
	}

}
