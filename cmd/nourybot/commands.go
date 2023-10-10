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
	go common.CommandUsed()

	go app.InitUser(message.User.Name, message.User.ID)

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

	case "catbox":
		go app.NewDownload("catbox", target, cmdParams[1], message)

	case "kappa":
		go app.NewDownload("kappa", target, cmdParams[1], message)

	case "yaf":
		go app.NewDownload("yaf", target, cmdParams[1], message)

	case "gofile":
		go app.NewDownload("gofile", target, cmdParams[1], message)

	case "osrs":
		reply = commands.OSRS(message.Message[7:len(message.Message)])

	case "preview":
		reply = commands.Preview(cmdParams[1])

	case "thumbnail":
		reply = commands.Preview(cmdParams[1])

	case "ffz":
		reply = commands.Ffz(cmdParams[1])

	case "ddg":
		reply = commands.DuckDuckGo(message.Message[6:len(message.Message)])

	case "youtube":
		reply = commands.Youtube(message.Message[10:len(message.Message)])

	case "godocs":
		reply = commands.Godocs(message.Message[9:len(message.Message)])

	case "google":
		reply = commands.Google(message.Message[9:len(message.Message)])

	case "duckduckgo":
		reply = commands.DuckDuckGo(message.Message[13:len(message.Message)])

	case "seventv":
		reply = commands.SevenTV(cmdParams[1])

	case "7tv":
		reply = commands.SevenTV(cmdParams[1])

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

	case "help":
		if msgLen > 1 {
			app.commandHelp(target, cmdParams[1], message.User.Name)
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

	default:
		r, err := app.GetCommand(target, commandName, userLevel)
		if err != nil {
			return
		}
		reply = r
	}
	if reply != "" {
		go app.Send(target, reply)
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
	"phonetic":   "Translates the input to the text equivalent on a phonetic russian keyboard layout. Layout and general functionality is the same as https://russian.typeit.org/",
	"ph":         "Translates the input to the text equivalent on a phonetic russian keyboard layout. Layout and general functionality is the same as https://russian.typeit.org/",
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
func (app *application) commandHelp(target, name, username string) {
	// Check if the `helpText` map has an entry for `name`. If it does return it's value entry
	// and send that as a reply.
	i, ok := helpText[name]
	if !ok {
		// If it doesn't check the database for a command with that `name`. If there is one
		// reply with that commands `help` entry.
		c, err := app.GetCommandHelp(name, username)
		if err != nil {
			app.Log.Infow("commandHelp: no such command found",
				"err", err)
			return
		}

		app.Send(target, c)
		return
	}

	app.Send(target, i)
}
