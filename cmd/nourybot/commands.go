package main

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/lyx0/nourybot/internal/commands"
	"github.com/lyx0/nourybot/internal/common"
)

// handleCommand takes in a twitch.PrivateMessage and then routes the message to
// the function that is responsible for each command and knows how to deal with it accordingly.
func (app *application) handleCommand(message twitch.PrivateMessage) {
	var reply string

	// Increments the counter how many commands have been used, called in the ping command.
	common.CommandUsed()

	// commandName is the actual name of the command without the prefix.
	// e.g. `()ping` would be `ping`.
	commandName := strings.ToLower(strings.SplitN(message.Message, " ", 3)[0][2:])

	// cmdParams are additional command parameters.
	// e.g. `()weather san antonio`
	// cmdParam[0] is `san` and cmdParam[1] = `antonio`.
	//
	// Since Twitch messages are at most 500 characters I use a
	// maximum count of 500+10 just to be safe.
	// https://discuss.dev.twitch.tv/t/missing-client-side-message-length-check/21316
	cmdParams := strings.SplitN(message.Message, " ", 500)

	// msgLen is the amount of words in a message without the prefix.
	// Useful to check if enough cmdParams are provided.
	msgLen := len(strings.SplitN(message.Message, " ", -2))

	userLevel := app.GetUserLevel(message.User.ID)
	// target is the channelname the message originated from and
	// where the TwitchClient should send the response
	target := message.Channel
	app.Log.Infow("Command received",
		// "message", message, // Pretty taxing
		"message.Message", message.Message,
		"message.Channel", target,
		"commandName", commandName,
		"cmdParams", cmdParams,
		"msgLen", msgLen,
		"userLevel", userLevel,
	)

	// A `commandName` is every message starting with `()`.
	// Hardcoded commands have a priority over database commands.
	// Switch over the commandName and see if there is a hardcoded case for it.
	// If there was no switch case satisfied, query the database if there is
	// a data.CommandModel.Name equal to the `commandName`
	// If there is return the data.CommandModel.Text entry.
	// Otherwise we ignore the message.
	switch commandName {
	case "":
		if msgLen == 1 {
			reply = "xd"
		}

	case "bttv":
		if msgLen < 2 {
			reply = "Not enough arguments provided. Usage: ()bttv <emote name>"
		} else {
			reply = commands.Bttv(cmdParams[1])
		}

		// Coinflip
	case "coin":
		reply = commands.Coinflip()
	case "coinflip":
		reply = commands.Coinflip()
	case "cf":
		reply = commands.Coinflip()

		// ()currency <amount> <input currency> to <output currency>
	case "currency":
		if msgLen < 4 {
			reply = "Not enough arguments provided. Usage: ()currency 10 USD to EUR"
		} else {
			reply, _ = commands.Currency(cmdParams[1], cmdParams[2], cmdParams[4])
		}

	case "dl":
		go commands.Download(target, cmdParams[1], fileUploaderURL, app.TwitchClient, app.Log)

	case "mail":
		app.SendEmail("Test command used!", "This is an email test")

	case "lastfm":
		if msgLen == 1 {
			reply = app.UserCheckLastFM(message)
		} else {
			// Default to first argument supplied being the name
			// of the user to look up recently played.
			reply = commands.LastFmUserRecent(target, cmdParams[1])
		}

	case "nourybot":
		reply = "Lidl Twitch bot made by @nourylul. Prefix: ()"

	case "phonetic":
		if msgLen == 1 {
			reply = "Not enough arguments provided. Usage: ()phonetic <text to translate>"
		} else {
			reply, _ = commands.Phonetic(message.Message[11:len(message.Message)])
		}
	case "ping":
		reply = commands.Ping()
		// ()bttv <emote name>

		// ()weather <location>
	case "weather":
		if msgLen == 1 {
			app.UserCheckWeather(message)
		} else if msgLen < 2 {
			reply = "Not enough arguments provided."
		} else {
			reply, _ = commands.Weather(message.Message[10:len(message.Message)])
		}

		// Xkcd
	// Random Xkcd
	case "rxkcd":
		reply, _ = commands.RandomXkcd()
	case "randomxkcd":
		reply, _ = commands.RandomXkcd()
	// Latest Xkcd
	case "xkcd":
		reply, _ = commands.Xkcd()

	case "timer":
		switch cmdParams[1] {
		case "add":
			app.AddTimer(cmdParams[2], cmdParams[3], message)
		case "edit":
			app.EditTimer(cmdParams[2], cmdParams[3], message)
		case "delete":
			app.DeleteTimer(cmdParams[2], message)
		case "list":
			reply = app.ListTimers()
		}

	case "debug":
		switch cmdParams[1] {
		case "user":
			if userLevel >= 250 {
				app.DebugUser(cmdParams[2], message)
			}
		case "command":
			if userLevel >= 250 {
				app.DebugCommand(cmdParams[2], message)
			}
		}

	case "command":
		switch cmdParams[1] {
		case "add":
			app.AddCommand(cmdParams[2], message)
		case "delete":
			app.DeleteCommand(cmdParams[2], message)
		case "edit":
			switch cmdParams[2] {
			case "level":
				app.EditCommandLevel(cmdParams[3], cmdParams[4], message)
			case "category":
				app.EditCommandCategory(cmdParams[3], cmdParams[4], message)
			}
		}

	case "set":
		switch cmdParams[1] {
		case "lastfm":
			app.SetUserLastFM(cmdParams[2], message)
		case "location":
			app.SetUserLocation(message)
		}

	case "user":
		switch cmdParams[1] {
		case "edit":
			switch cmdParams[2] {
			case "level":
				app.EditUserLevel(cmdParams[3], cmdParams[4], message)
			}
		}
	}
	if reply != "" {
		go app.Send(target, reply)
		return
	}
}
