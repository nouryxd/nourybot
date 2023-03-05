package main

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/commands"
	"github.com/lyx0/nourybot/internal/common"
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
	// where the TwitchClient should send the response
	target := message.Channel

	// Userlevel is the level set for a user in the database.
	// It is NOT same as twitch user/mod.
	// 1000 = admin
	// 500 = mod
	// 250 = vip
	// 100 = normal
	// If the level returned is 0 then the user was not found in the database.
	userLevel := app.GetUserLevel(message.User.ID)

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
	// a data.CommandModel.Name equal to the `commandName`
	// If there is return the data.CommandModel.Text entry.
	// Otherwise we ignore the message.
	switch commandName {
	case "":
		if msgLen == 1 {
			common.Send(target, "xd", app.TwitchClient)
			return
		}

	case "nourybot":
		common.Send(target, "Lidl Twitch bot made by @nourylul. Prefix: ()", app.TwitchClient)
		return

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
	case "lastfm":
		if msgLen == 1 {
			app.UserCheckLastFM(message)
			return
		} else if cmdParams[1] == "artist" && cmdParams[2] == "top" {
			commands.LastFmArtistTop(target, message, app.TwitchClient)
			return
		} else {
			// Default to first argument supplied being the name
			// of the user to look up recently played.
			commands.LastFmUserRecent(target, cmdParams[1], app.TwitchClient)
			return
		}

	case "help":
		if msgLen == 1 {
			common.Send(target, "Provides information for a given command. Usage: ()help <commandname>", app.TwitchClient)
			return
		} else {
			app.commandHelp(target, cmdParams[1], message.User.Name)
			return
		}

	// ()ping
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

	case "set":
		if msgLen < 3 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else if cmdParams[1] == "lastfm" {
			app.SetUserLastFM(cmdParams[2], message)
			//app.SetLastFMUser(cmdParams[2], message)
			return
		} else if cmdParams[1] == "location" {
			app.SetUserLocation(message)
			return
		} else {
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

	// ()thumbnail <live channel>
	case "thumbnail":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()thumbnail <username>", app.TwitchClient)
			return
		} else {
			commands.Preview(target, cmdParams[1], app.TwitchClient)
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

	// ()weather <location>
	case "weather":
		if msgLen == 1 {
			// Default to first argument supplied being the name
			// of the user to look up recently played.
			app.UserCheckWeather(message)
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			commands.Weather(target, message.Message[10:len(message.Message)], app.TwitchClient)
			return
		}

	// Xkcd
	// Random Xkcd
	case "rxkcd":
		commands.RandomXkcd(target, app.TwitchClient)
		return
	case "randomxkcd":
		commands.RandomXkcd(target, app.TwitchClient)
		return
	// Latest Xkcd
	case "xkcd":
		commands.Xkcd(target, app.TwitchClient)
		return

	// Commands with permission level or database from here on

	//#################
	// 250 - VIP only
	//#################
	// ()debug user <username>
	// ()debug command <command name>
	case "debug":
		if userLevel < 250 {
			return
		} else if msgLen < 3 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else if cmdParams[1] == "user" {
			app.DebugUser(cmdParams[2], message)
			return
		} else if cmdParams[1] == "command" {
			app.DebugCommand(cmdParams[2], message)
			return
		} else {
			return
		}

	// ()echo <message>
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
	// 1000 - Admin only
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
	case "addtimer":
		if userLevel < 1000 {
			return
		} else if msgLen < 4 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()addtimer gfuel 5m sponsor XD xD
			app.AddTimer(cmdParams[1], message)
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
		// ()edittimer testname 10m test text xd
	case "edittimer":
		if userLevel < 1000 { // Limit to myself for now.
			return
		} else if msgLen < 4 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()edituser level nourylul 1000
			app.EditTimer(cmdParams[1], message)
		}
	case "editcommand": // ()editcommand level dankwave 1000
		if userLevel < 1000 {
			return
		} else if msgLen < 4 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else if cmdParams[1] == "level" {
			app.EditCommandLevel(cmdParams[2], cmdParams[3], message)
			return
		} else if cmdParams[1] == "category" {
			app.EditCommandCategory(cmdParams[2], cmdParams[3], message)
			return
		} else if cmdParams[1] == "help" {
			app.EditCommandHelp(cmdParams[2], message)
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
	case "deletetimer":
		if userLevel < 1000 { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()deletetimer dank
			app.DeleteTimer(cmdParams[1], message)
			return
		}

	case "asd":
		app.Logger.Info(app.Scheduler.Entries())
		return

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

// Map of known commands with their help texts.
var helpText = map[string]string{
	"bttv":       "Returns the search URL for a given BTTV emote. Example usage: ()bttv <emote name>",
	"coin":       "Flips a coin! Aliases: coinflip, coin, cf",
	"cf":         "Flips a coin! Aliases: coinflip, coin, cf",
	"coinflip":   "Flips a coin! Aliases: coinflip, coin, cf",
	"currency":   "Returns the exchange rate for two currencies. Only three letter abbreviations are supported ( List of supported currencies: https://decapi.me/misc/currency?list ). Example usage: ()currency 10 USD to EUR",
	"ffz":        "Returns the search URL for a given FFZ emote. Example usage: ()ffz <emote name>",
	"followage":  "Returns how long a given user has been following a channel. Example usage: ()followage <channel> <username>",
	"firstline":  "Returns the first message a user has sent in a given channel. Aliases: firstline, fl. Example usage: ()firstline <channel> <username>",
	"fl":         "Returns the first message a user has sent in a given channel. Aliases: firstline, fl. Example usage: ()fl <channel> <username>",
	"help":       "Returns more information about a command and its usage. 4Head Example usage: ()help <command name>",
	"ping":       "Hopefully returns a Pong! monkaS",
	"preview":    "Returns a link to an (almost) live screenshot of a live channel. Alias: preview, thumbnail. Example usage: ()preview <channel>",
	"thumbnail":  "Returns a link to an (almost) live screenshot of a live channel. Alias: preview, thumbnail. Example usage: ()thumbnail <channel>",
	"tweet":      "Returns the latest tweet for a provided user. Example usage: ()tweet <username>",
	"seventv":    "Returns the search URL for a given SevenTV emote. Aliases: seventv, 7tv. Example usage: ()seventv FeelsDankMan",
	"7tv":        "Returns the search URL for a given SevenTV emote. Aliases: seventv, 7tv. Example usage: ()7tv FeelsDankMan",
	"weather":    "Returns the weather for a given location. Example usage: ()weather Vilnius",
	"randomxkcd": "Returns a link to a random xkcd comic. Alises: randomxkcd, rxkcd. Example usage: ()randomxkcd",
	"rxkcd":      "Returns a link to a random xkcd comic. Alises: randomxkcd, rxkcd. Example usage: ()rxkcd",
	"xkcd":       "Returns a link to the latest xkcd comic. Example usage: ()xkcd",
}

// Help checks if a help text for a given command exists and replies with it.
func (app *Application) commandHelp(target, name, username string) {
	// Check if the `helpText` map has an entry for `name`. If it does return it's value entry
	// and send that as a reply.
	i, ok := helpText[name]
	if !ok {
		// If it doesn't check the database for a command with that `name`. If there is one
		// reply with that commands `help` entry.
		c, err := app.GetCommandHelp(name, username)
		if err != nil {
			app.Logger.Infow("commandHelp: no such command found",
				"err", err)
			return
		}

		reply := fmt.Sprint(c)
		common.Send(target, reply, app.TwitchClient)
		return
	}

	reply := fmt.Sprint(i)
	common.Send(target, reply, app.TwitchClient)
}
