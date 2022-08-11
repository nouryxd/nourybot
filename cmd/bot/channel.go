package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/commands/decapi"
	"github.com/lyx0/nourybot/pkg/common"
)

// AddChannel takes in a channel name, then queries `GetIdByLogin` for the
// channels ID and inserts both the name and id value into the database.
// If there was no error thrown the `app.TwitchClient` joins the channel afterwards.
func (app *Application) AddChannel(login string, message twitch.PrivateMessage) {
	userId, err := decapi.GetIdByLogin(login)
	if err != nil {
		app.Logger.Error(err)
		return
	}

	// Initialize a new channel struct holding the values that will be
	// passed into the `app.Models.Channels.Insert()` method.
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

// GetAllChannels() queries the database and lists all channels.
// Only used for debug/information purposes.
func (app *Application) GetAllChannels() {
	channel, err := app.Models.Channels.GetAll()
	if err != nil {
		app.Logger.Error(err)
		return
	}
	app.Logger.Infow("All channels:",
		"channel", channel)
}

// DeleteChannel queries the database for a channel name and if it exists
// deletes the channel and makes the bot depart said channel.
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

// InitialJoin is called on startup and gets all the channels from the database
// so that the `app.TwitchClient` then joins each.
func (app *Application) InitialJoin() {
	// GetJoinable returns a `[]string` containing the channel names.
	channel, err := app.Models.Channels.GetJoinable()
	if err != nil {
		app.Logger.Error(err)
		return
	}

	// Iterate over the `[]string` and join each.
	for _, v := range channel {
		app.TwitchClient.Join(v)
		app.Logger.Infow("Joining channel:",
			"channel", v)
	}
}
