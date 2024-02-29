package main

import (
	"github.com/nicklaw5/helix/v2"
)

// getUserID returns the Twitch userID for given twitch username.
func (app *application) getUserID(username string) (string, error) {
	resp, err := app.HelixClient.GetUsers(&helix.UsersParams{
		Logins: []string{username},
	})
	if err != nil {
		app.Log.Error(err)
		return "", err
	}

	return resp.Data.Users[0].ID, nil
}

// userIDtoUsername returns the Twitch username for a given twitch user ID.
func (app *application) userIDtoUsername(channelID string) (string, error) {
	resp, err := app.HelixClient.GetUsers(&helix.UsersParams{
		IDs: []string{channelID},
	})
	if err != nil {
		app.Log.Error(err)
		return "", err
	}

	return resp.Data.Users[0].Login, nil
}

// getChannelTitle queries the Twitch API for the current Title for
// the provided Twitch channel username.
func (app *application) getChannelTitle(channelID string) string {
	resp, err := app.HelixClient.GetChannelInformation(&helix.GetChannelInformationParams{
		BroadcasterID: channelID,
	})
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	return resp.Data.Channels[0].Title
}

// getChannelTitleByUsername returns the current title of the provided
// Twitch channel username.
func (app *application) getChannelTitleByUsername(username string) string {
	channelID, err := app.getUserID(username)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	reply := app.getChannelTitle(channelID)
	return reply
}

func (app *application) getChannelGame(channelId string) string {
	resp, err := app.HelixClient.GetChannelInformation(&helix.GetChannelInformationParams{
		BroadcasterID: channelId,
	})
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	return resp.Data.Channels[0].GameName
}

func (app *application) getChannelGameByUsername(username string) string {
	channelID, err := app.getUserID(username)
	if err != nil {
		app.Log.Error(err)
		return "Something went wrong FeelsBadMan"
	}

	reply := app.getChannelGame(channelID)
	return reply
}
