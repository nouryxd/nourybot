package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

// AddChannel looks up the userID of the provided login and tries to insert
// the Twitch login and userID into the database. If there is no error thrown the
// TwitchClient joins the channel afterwards.
func (app *application) AddChannel(login string, message twitch.PrivateMessage) {
	userID, err := app.getUserID(login)
	if err != nil {
		app.Log.Error(err)
		return
	}

	err = app.Models.Channels.Insert(login, userID)
	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		app.Send(message.Channel, reply, message)
		return
	} else {
		app.TwitchClient.Join(login)
		reply := fmt.Sprintf("Added channel %s", login)
		app.Send(message.Channel, reply, message)
		return
	}
}

// GetAllChannels() queries the database and lists all channels.
func (app *application) GetAllChannels() {
	channel, err := app.Models.Channels.GetAll()
	if err != nil {
		app.Log.Error(err)
		return
	}
	app.Log.Infow("All channels:",
		"channel", channel)
}

// DeleteChannel queries the database for a channel name and if it exists
// deletes the channel and makes the bot depart said channel.
func (app *application) DeleteChannel(login string, message twitch.PrivateMessage) {
	err := app.Models.Channels.Delete(login)
	if err != nil {
		app.Send(message.Channel, "Something went wrong FeelsBadMan", message)
		app.Log.Error(err)
		return
	}

	app.TwitchClient.Depart(login)

	reply := fmt.Sprintf("Deleted channel %s", login)
	app.Send(message.Channel, reply, message)
}

// InitialJoin is called on startup and queries the database for a list of
// channels which the TwitchClient then joins.
func (app *application) InitialJoin() {
	// GetJoinable returns a slice of channel names.
	channel, err := app.Models.Channels.GetJoinable()
	if err != nil {
		app.Log.Error(err)
		return
	}

	// Iterate over the slice of channels and join each.
	for _, v := range channel {
		app.TwitchClient.Join(v)
		app.Log.Infow("Joining channel",
			"channel", v)
	}
}
