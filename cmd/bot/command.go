package main

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/common"
)

// handleCommand takes in a twitch.PrivateMessage and then routes the message to
// the function that is responsible for each command and knows how to deal with it accordingly.
func (app *Application) handleCommand(message twitch.PrivateMessage) {

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

	// target is the channelname the message originated from and
	// where we are responding.
	target := message.Channel

	// Userlevel is the level set for a user in the database.
	// It is NOT same as twitch user/mod.
	// 1000 = admin
	// 500 = mod
	// 250 = vip
	// 100 = normal
	// If the level returned is 0 then the user was not found in the database.
	userLevel := app.GetUserLevel(message.User.Name)

	// Left in for debugging purposes.
	app.Logger.Infow("Command received",
		// "message", message, // Pretty taxing
		"message.Message", message.Message,
		"message.Channel", target,
		"userLevel", userLevel,
		"commandName", commandName,
		"cmdParams", cmdParams,
		"msgLen", msgLen,
	)

	// A `commandName` is every message starting with `()`.
	// Hardcoded commands have a priority over database commands.
	// Switch over the commandName and see if there is a hardcoded case for it.
	// If there was no switch case satisfied, query the database if there is
	// a `data.CommandModel.Name` equal to the `commandName`
	// If there is return the `data.CommandModel.Text` entry.
	// Otherwise we ignore the message.
	switch commandName {
	case "":
		if msgLen == 1 {
			common.Send(target, "xd", app.TwitchClient)
			return
		}

	// ()bttv <emote name>
	case "bttv":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()bttv <emote name>", app.TwitchClient)
			return
		} else {
			commands.Bttv(target, cmdParams[1], app.TwitchClient)
			return
		}

	// Coinflip
	case "coin":
		commands.Coinflip(target, app.TwitchClient)
		return
	case "coinflip":
		commands.Coinflip(target, app.TwitchClient)
		return
	case "cf":
		commands.Coinflip(target, app.TwitchClient)
		return

	// ()currency <amount> <input currency> to <output currency>
	case "currency":
		if msgLen < 4 {
			common.Send(target, "Not enough arguments provided. Usage: ()currency 10 USD to EUR", app.TwitchClient)
			return
		}
		commands.Currency(target, cmdParams[1], cmdParams[2], cmdParams[4], app.TwitchClient)
		return

	// ()ffz <emote name>
	case "ffz":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()ffz <emote name>", app.TwitchClient)
			return
		} else {
			commands.Ffz(target, cmdParams[1], app.TwitchClient)
			return
		}

	// ()followage <channel> <username>
	case "followage":
		if msgLen == 1 { // ()followage
			commands.Followage(target, target, message.User.Name, app.TwitchClient)
			return
		} else if msgLen == 2 { // ()followage forsen
			commands.Followage(target, target, cmdParams[1], app.TwitchClient)
			return
		} else { // ()followage forsen pajlada
			commands.Followage(target, cmdParams[1], cmdParams[2], app.TwitchClient)
			return
		}

	// First Line
	// ()firstline <channel> <username>
	case "firstline":
		if msgLen == 1 {
			common.Send(target, "Usage: ()firstline <channel> <user>", app.TwitchClient)
			return
		} else if msgLen == 2 {
			commands.FirstLine(target, target, cmdParams[1], app.TwitchClient)
			return
		} else {
			commands.FirstLine(target, cmdParams[1], cmdParams[2], app.TwitchClient)
			return
		}
	// ()fl <channel> <username>
	case "fl":
		if msgLen == 1 {
			common.Send(target, "Usage: ()firstline <channel> <user>", app.TwitchClient)
			return
		} else if msgLen == 2 {
			commands.FirstLine(target, target, cmdParams[1], app.TwitchClient)
			return
		} else {
			commands.FirstLine(target, cmdParams[1], cmdParams[2], app.TwitchClient)
			return
		}

	case "ping":
		commands.Ping(target, app.TwitchClient)
		return

	// Thumbnail
	// ()preview <live channel>
	case "preview":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()preview <username>", app.TwitchClient)
			return
		} else {
			commands.Preview(target, cmdParams[1], app.TwitchClient)
			return
		}
	// ()thumbnail <live channel>
	case "thumbnail":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()thumbnail <username>", app.TwitchClient)
			return
		} else {
			commands.Preview(target, cmdParams[1], app.TwitchClient)
			return
		}

	// SevenTV
	// ()seventv <emote name>
	case "seventv":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()seventv <emote name>", app.TwitchClient)
			return
		} else {
			commands.Seventv(target, cmdParams[1], app.TwitchClient)
			return
		}
	// ()7tv <emote name>
	case "7tv":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()seventv <emote name>", app.TwitchClient)
			return
		} else {
			commands.Seventv(target, cmdParams[1], app.TwitchClient)
			return
		}

	// ()tweet <username>
	case "tweet":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()tweet <username>", app.TwitchClient)
			return
		} else {
			commands.Tweet(target, cmdParams[1], app.TwitchClient)
			return
		}

	// Xkcd
	case "rxkcd":
		commands.RandomXkcd(target, app.TwitchClient)
		return
	case "randomxkcd":
		commands.RandomXkcd(target, app.TwitchClient)
		return
	case "xkcd":
		commands.Xkcd(target, app.TwitchClient)
		return

	// Commands with permission level or database from here on

	//#################
	// 250 - VIP only
	//#################
	case "debug":
		if userLevel < 250 {
			return
		} else if msgLen < 3 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else if cmdParams[1] == "user" {
			app.DebugUser(cmdParams[2], message)
			return
		} else {
			return
		}

	case "echo":
		if userLevel < 250 { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			commands.Echo(target, message.Message[7:len(message.Message)], app.TwitchClient)
			return
		}

	//###################
	// 1000- Admin only
	//###################

	// #####
	// Add
	// #####
	case "addchannel":
		if userLevel < 1000 {
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()addchannel noemience
			app.AddChannel(cmdParams[1], message)
			return
		}
	case "addcommand":
		if userLevel < 1000 {
			return
		} else if msgLen < 3 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()addcommand dank FeelsDankMan xD
			app.AddCommand(cmdParams[1], message)
			return
		}
	case "adduser":
		if userLevel < 1000 {
			return
		} else if msgLen < 3 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()adduser nourylul 1000
			app.AddUser(cmdParams[1], cmdParams[2], message)
			return
		}

	// ######
	// Edit
	// ######
	case "edituser":
		if userLevel < 1000 { // Limit to myself for now.
			return
		} else if msgLen < 4 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else if cmdParams[1] == "level" {
			// ()edituser level nourylul 1000
			app.EditUserLevel(cmdParams[2], cmdParams[3], message)
			return
		} else {
			return
		}
	case "editcommand": // ()editcommand level nourylul 1000
		if userLevel < 1000 {
			return
		} else if msgLen < 4 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else if cmdParams[1] == "level" {
			app.EditCommandLevel(cmdParams[2], cmdParams[3], message)
			return
		} else {
			return
		}

	// ########
	// Delete
	// ########
	case "deletechannel":
		if userLevel < 1000 { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()deletechannel noemience
			app.DeleteChannel(cmdParams[1], message)
			return
		}
	case "deletecommand":
		if userLevel < 1000 { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()deletecommand dank
			app.DeleteCommand(cmdParams[1], message)
			return
		}
	case "deleteuser":
		if userLevel < 1000 { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()deleteuser noemience
			app.DeleteUser(cmdParams[1], message)
			return
		}

	case "bttvemotes":
		if userLevel < 1000 {
			commands.Bttvemotes(target, app.TwitchClient)
			return
		} else {
			return
		}

	case "ffzemotes":
		if userLevel < 1000 {
			commands.Ffzemotes(target, app.TwitchClient)
			return
		} else {
			return
		}

	// ##################
	// Database Command
	// Check if the commandName exists as the "name" of a command in the database.
	// if it doesnt then ignore it.
	// ##################
	default:
		reply, err := app.GetCommand(commandName, message.User.Name)
		if err != nil {
			return
		}
		common.SendNoLimit(message.Channel, reply, app.TwitchClient)
		return
	}
}
