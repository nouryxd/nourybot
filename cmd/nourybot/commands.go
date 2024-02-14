package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/lyx0/nourybot/internal/commands"
	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/pkg/ivr"
	"github.com/lyx0/nourybot/pkg/lastfm"
	"github.com/lyx0/nourybot/pkg/owm"
)

// handleCommand takes in a twitch.PrivateMessage and then routes the message to
// the function that is responsible for each command and knows how to deal with it accordingly.
func (app *application) handleCommand(message twitch.PrivateMessage) {
	var reply string

	if message.Channel == "forsen" {
		return
	}

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

	userLevel := app.GetUserLevel(message)

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

	go app.LogCommand(message, commandName, userLevel)
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
		// --------------------------------
		// pleb commands
		// --------------------------------
	case "bttv":
		if msgLen < 2 {
			reply = "Not enough arguments provided. Usage: ()bttv <emote name>"
		} else {
			reply = commands.Bttv(cmdParams[1])
		}

	case "betterttv":
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

	case "osrs":
		reply = commands.OSRS(message.Message[7:len(message.Message)])

	case "preview":
		reply = commands.Preview(cmdParams[1])

	case "thumbnail":
		reply = commands.Preview(cmdParams[1])

	case "ffz":
		reply = commands.Ffz(cmdParams[1])

	case "frankerfacez":
		reply = commands.Ffz(cmdParams[1])

	case "notify":
		switch cmdParams[1] {
		case "live":
			if userLevel >= 100 {
				reply = app.createLiveSubscription(target, cmdParams[2])
			}
		case "offline":
			if userLevel >= 100 {
				reply = app.createOfflineSubscription(target, cmdParams[2])
			}
		}
	case "unnotify":
		switch cmdParams[1] {
		case "live":
			if userLevel >= 100 {
				reply = app.deleteLiveSubscription(target, cmdParams[2])
			}
		case "offline":
			if userLevel >= 100 {
				reply = app.deleteOfflineSubscription(target, cmdParams[2])
			}
		}

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

	case "lastfm":
		if msgLen == 1 {
			reply = app.UserCheckLastFM(message)
		} else {
			// Default to first argument supplied being the name
			// of the user to look up recently played.
			reply = lastfm.LastFmUserRecent(target, cmdParams[1])
		}

	case "help":
		if msgLen > 1 {
			app.commandHelp(target, cmdParams[1], message.User.Name, message)
		}

	case "nourybot":
		reply = "Lidl Twitch bot made by @nouryxd. Prefix: ()"

	case "predb":
		switch cmdParams[1] {
		case "latest":
			if userLevel >= 100 {
				reply = app.PreDBLatest()
			}
		case "search":
			if userLevel >= 100 && len(message.Message) > 16 {
				reply = app.PreDBSearch(message.Message[15:len(message.Message)])
			}
		case "group":
			if userLevel >= 100 && len(message.Message) > 15 {
				reply = app.PreDBGroup(message.Message[14:len(message.Message)])
			}
		}

	case "phonetic":
		if msgLen == 1 {
			reply = "Not enough arguments provided. Usage: ()phonetic <text to translate>"
		} else {
			reply, _ = commands.Phonetic(message.Message[11:len(message.Message)])
		}
	case "ping":
		reply = commands.Ping(app.Environment)
		// ()bttv <emote name>

		// ()weather <location>
	case "weather":
		if msgLen == 1 {
			app.UserCheckWeather(message)
		} else if msgLen < 2 {
			reply = "Not enough arguments provided."
		} else {
			reply, _ = owm.Weather(message.Message[10:len(message.Message)])
		}

	case "rxkcd":
		reply, _ = commands.RandomXkcd()

	case "randomxkcd":
		reply, _ = commands.RandomXkcd()

	case "xkcd":
		reply, _ = commands.Xkcd()

	case "uid":
		reply = ivr.IDByUsernameReply(cmdParams[1])

	case "userid":
		reply = ivr.IDByUsernameReply(cmdParams[1])

	case "commands":
		reply = app.ListChannelCommands(message.Channel)

	case "timers":
		reply = fmt.Sprintf("https://bot.noury.li/timer/%s", message.Channel)

	case "conv":
		if userLevel >= 100 {
			app.ConvertToMP4(cmdParams[1], message)
		}

	case "meme":
		if userLevel >= 100 {
			app.ConvertAndSave(cmdParams[1], cmdParams[2], message)
		}

	case "set":
		switch cmdParams[1] {
		case "lastfm":
			app.SetUserLastFM(cmdParams[2], message)
		case "location":
			app.SetUserLocation(message)
		}

		// --------------------------------
		// 100 user level
		// trusted
		// vip
		// --------------------------------
	case "wa":
		if userLevel >= 100 {
			reply = commands.WolframAlphaQuery(message.Message[5:len(message.Message)], app.Config.wolframAlphaAppID)
		}

	case "query":
		if userLevel >= 100 {
			reply = commands.WolframAlphaQuery(message.Message[8:len(message.Message)], app.Config.wolframAlphaAppID)
		}
	case "debug":
		switch cmdParams[1] {
		case "user":
			if userLevel >= 100 {
				app.DebugUser(cmdParams[2], message)
			}
		case "command":
			if userLevel >= 100 {
				app.DebugCommand(cmdParams[2], message)
			}
		case "env":
			if userLevel >= 100 {
				reply = app.Environment
			}
		case "timers":
			if userLevel >= 100 {
				app.DebugChannelTimers(message.Channel)
			}
		}

		// --------------------------------
		// 250 user level
		// twitch mod/broadcaster
		// --------------------------------
		// empty for now

		// --------------------------------
		// 420 user level
		// dank
		// --------------------------------
	case "catbox":
		if userLevel >= 420 {
			go app.NewDownload("catbox", target, cmdParams[1], message)
		}

	case "kappa":
		if userLevel >= 420 {
			go app.NewDownload("kappa", target, cmdParams[1], message)
		}

	case "yaf":
		if userLevel >= 420 {
			go app.NewDownload("yaf", target, cmdParams[1], message)
		}

	case "gofile":
		if userLevel >= 420 {
			go app.NewDownload("gofile", target, cmdParams[1], message)
		}

		//----------------------------------
		// 500 User Level
		// trusted
		//---------------------------------
	case "timer":
		switch cmdParams[1] {
		case "add":
			if userLevel >= 500 {
				app.AddTimer(cmdParams[2], cmdParams[3], message)
			}
		case "delete":
			if userLevel >= 500 {
				app.DeleteTimer(cmdParams[2], message)
			}
		case "edit":
			if userLevel >= 500 {
				app.EditTimer(cmdParams[2], cmdParams[3], message)
			}
		}

	case "command":
		switch cmdParams[1] {
		case "add":
			if userLevel >= 500 {
				app.AddCommand(cmdParams[2], message)
			}
		case "delete":
			if userLevel >= 500 {
				app.DeleteCommand(cmdParams[2], message)
			}
		case "list":
			if userLevel >= 500 {
				reply = app.ListCommands()
			}
		case "edit":
			switch cmdParams[2] {
			case "level":
				if userLevel >= 500 {
					app.EditCommandLevel(cmdParams[3], cmdParams[4], message)
				}
			}
		}

		//------------------------------------
		// 1000 User Level
		// Admin
		//------------------------------------
	case "join":
		if userLevel >= 1000 {
			go app.AddChannel(cmdParams[1], message)
		}

	case "part":
		if userLevel >= 1000 {
			go app.DeleteChannel(cmdParams[1], message)
		}

	case "user":
		switch cmdParams[1] {
		case "edit":
			switch cmdParams[2] {
			case "level":
				if userLevel >= 1000 {
					app.EditUserLevel(cmdParams[3], cmdParams[4], message)
				}
			}
		case "set":
			switch cmdParams[2] {
			case "level":
				if userLevel >= 1000 {
					app.EditUserLevel(cmdParams[3], cmdParams[4], message)
				}
			}
		}

	default:
		r, err := app.GetCommand(target, commandName, userLevel)
		if err != nil {
			return
		}
		go app.SendNoBanphrase(target, r)
		return
	}
	if reply != "" {
		go app.Send(target, reply, message)
		return
	}
}

type command struct {
	Name        string
	Alias       []string
	Description string
	Level       string
	Usage       string
}

// Optional is []
// Required is < >
var helpText = map[string]command{
	"bttv": {
		Name:        "bttv",
		Alias:       []string{"bttv", "betterttv"},
		Description: "Returns the search URL for a given BTTV emote.",
		Level:       "0",
		Usage:       "()bttv <emote name>",
	},
	"catbox": {
		Name:        "catbox",
		Alias:       nil,
		Description: "Downloads the video of a given link with yt-dlp and then uploads the video to catbox.moe.",
		Level:       "420",
		Usage:       "()catbox <link>",
	},
	"coin": {
		Name:        "coin",
		Alias:       []string{"coin", "coinflip", "cf"},
		Description: "Flip a coin.",
		Level:       "0",
		Usage:       "()coin",
	},
	"command add": {
		Name:        "command add",
		Alias:       nil,
		Description: "Adds a channel command to the database.",
		Level:       "250",
		Usage:       "()add command <command name> <command text>",
	},
	"command edit level": {
		Name:        "command edit level",
		Alias:       nil,
		Description: "Edits the required level of a channel command with the given name.",
		Level:       "250",
		Usage:       "()command edit level <command name> <new command level>",
	},
	"command delete": {
		Name:        "command delete",
		Alias:       nil,
		Description: "Deletes the channel command with the given name.",
		Level:       "250",
		Usage:       "()command delete <command name>",
	},
	"commands": {
		Name:        "commands",
		Alias:       nil,
		Description: "Returns a link to the commands in the channel.",
		Level:       "0",
		Usage:       "()commands",
	},
	"currency": {
		Name:        "currency",
		Alias:       []string{"currency", "money"},
		Description: "Returns the exchange rate for two currencies. Only three letter abbreviations are supported ( List of supported currencies: https://decapi.me/misc/currency?list ).",
		Level:       "0",
		Usage:       "()currency <curr> to <curr>",
	},
	"debug env": {
		Name:        "debug env",
		Alias:       nil,
		Description: "Returns the environment currently running in.",
		Level:       "100",
		Usage:       "()debug env",
	},
	"debug user": {
		Name:        "debug user",
		Alias:       nil,
		Description: "Returns additional information about a user.",
		Level:       "100",
		Usage:       "()debug user <username>",
	},
	"debug command": {
		Name:        "debug command",
		Alias:       nil,
		Description: "Returns additional informations about a command.",
		Level:       "100",
		Usage:       "()debug command <command name>",
	},
	"debug timers": {
		Name:        "debug timers",
		Alias:       nil,
		Description: "Returns a list of timers currently running in the channel with additional informations.",
		Level:       "100",
		Usage:       "()debug timers",
	},
	"duckduckgo": {
		Name:        "duckduckgo",
		Alias:       []string{"duckduckgo", "ddg"},
		Description: "Returns the duckduckgo search URL for a given query.",
		Level:       "0",
		Usage:       "()duckduckgo <query>",
	},
	"ffz": {
		Name:        "ffz",
		Alias:       []string{"ffz", "frankerfacez"},
		Description: "Returns the search URL for a given FFZ emote.",
		Level:       "0",
		Usage:       "()ffz <emote name>",
	},
	"firstline": {
		Name:        "firstline",
		Alias:       []string{"firstline", "fl"},
		Description: "Returns the first message a user has sent in a channel",
		Level:       "0",
		Usage:       "()firstline <channel> <username>",
	},
	"followage": {
		Name:        "followage",
		Alias:       []string{"followage", "fa"},
		Description: "Returns how long a user has been following a channel.",
		Level:       "0",
		Usage:       "()followage <channel> <username>",
	},
	"godocs": {
		Name:        "godocs",
		Alias:       nil,
		Description: "Returns the godocs.io search URL for a given query.",
		Level:       "0",
		Usage:       "()godoc <query>",
	},
	"gofile": {
		Name:        "gofile",
		Alias:       nil,
		Description: "Downloads the video of a given link with yt-dlp and then uploads the video to gofile.io.",
		Level:       "420",
		Usage:       "()gofile <link>",
	},
	"google": {
		Name:        "google",
		Alias:       nil,
		Description: "Returns the google search URL for a given query.",
		Level:       "0",
		Usage:       "()google <query>",
	},
	"help": {
		Name:        "help",
		Alias:       nil,
		Description: "Returns more information about a command 4Head.",
		Level:       "0",
		Usage:       "()help <command name>",
	},
	"join": {
		Name:        "join",
		Alias:       nil,
		Description: "Adds the bot to a given channel.",
		Level:       "1000",
		Usage:       "()join <channel name>",
	},
	"kappa": {
		Name:        "kappa",
		Alias:       nil,
		Description: "Downloads the video of a given link with yt-dlp and then uploads the video to kappa.lol.",
		Level:       "420",
		Usage:       "()kappa <link>",
	},
	"lastfm": {
		Name:        "lastfm",
		Alias:       nil,
		Description: `Look up the last played title for a user on last.fm. If you "$set lastfm" a last.fm username the command will use that username if no other username is specified.`,
		Level:       "0",
		Usage:       "()lastfm [username]",
	},
	"osrs": {
		Name:        "osrs",
		Alias:       nil,
		Description: "Returns the oldschool runescape wiki search URL for a given query.",
		Level:       "0",
		Usage:       "()osrs <query>",
	},
	"part": {
		Name:        "part",
		Alias:       nil,
		Description: "Removes the bot from a given channel.",
		Level:       "1000",
		Usage:       "()part <channel name>",
	},
	"phonetic": {
		Name:        "phonetic",
		Alias:       []string{"phonetic", "ph"},
		Description: "Translates the input to the character equivalent on a phonetic russian keyboard layout. Layout and general functionality is the same as on https://russian.typeit.org/",
		Level:       "0",
		Usage:       "()phonetic <input>",
	},
	"ping": {
		Name:        "ping",
		Alias:       nil,
		Description: "Hopefully returns a pong monkaS.",
		Level:       "0",
		Usage:       "()ping",
	},
	"predb search": {
		Name:        "predb search",
		Alias:       nil,
		Description: "Returns the last 100 predb.net search results for a given query.",
		Level:       "100",
		Usage:       "()predb search <query>",
	},
	"predb group": {
		Name:        "predb group",
		Alias:       nil,
		Description: "Returns the last 100 predb.net group results for a given release group.",
		Level:       "100",
		Usage:       "()predb group <group>",
	},
	"preview": {
		Name:        "preview",
		Alias:       []string{"preview", "thumbnail"},
		Description: "Returns a link to an (almost) live screenshot of a live channel.",
		Level:       "0",
		Usage:       "()thumbnail <channel>",
	},
	"randomxkcd": {
		Name:        "randomxkcd",
		Alias:       []string{"randomxkcd", "rxkcd"},
		Description: "Returns a link to a random xkcd comic.",
		Level:       "0",
		Usage:       "()randomxkcd",
	},
	"seventv": {
		Name:        "seventv",
		Alias:       []string{"seventv", "7tv"},
		Description: "Returns the search URL for a given SevenTV emote.",
		Level:       "0",
		Usage:       "()seventv <emote name>",
	},
	"set lastfm": {
		Name:        "set lastfm",
		Alias:       nil,
		Description: "Allows you to set a last.fm username that will be used for $lastfm lookups if no other username is specified.",
		Level:       "0",
		Usage:       "()set lastfm <username>",
	},
	"set location": {
		Name:        "set location",
		Alias:       nil,
		Description: "Allows you to set a location that will be used for $weather lookups if no other location is specified.",
		Level:       "0",
		Usage:       "()set location <location>",
	},
	"timer add": {
		Name:        "timer add",
		Alias:       nil,
		Description: "Adds a new timer to the channel.",
		Level:       "500",
		Usage:       "()timer add <name> <repeat> <text>",
	},
	"timer delete": {
		Name:        "timer delete",
		Alias:       nil,
		Description: "Deletes a timer from the channel.",
		Level:       "500",
		Usage:       "()timer delete <name>",
	},
	"timer edit": {
		Name:        "timer edit",
		Alias:       nil,
		Description: "Edits a timer from the channel.",
		Level:       "500",
		Usage:       "()timer edit <name> <repeat> <text>",
	},
	"timers": {
		Name:        "timers",
		Alias:       nil,
		Description: "Returns a link to the currently active timers in the channel.",
		Level:       "0",
		Usage:       "()timers",
	},
	"user edit level": {
		Name:        "user edit level",
		Alias:       []string{"user set level"},
		Description: "Edits the user level for a given username.",
		Level:       "1000",
		Usage:       "()user edit level <username> <level>",
	},
	"uid": {
		Name:        "uid",
		Alias:       []string{"uid", "userid"},
		Description: "Returns the Twitch user ID for a given username.",
		Level:       "0",
		Usage:       "()uid <username>",
	},
	"wa": {
		Name:        "wa",
		Alias:       []string{"wa", "query"},
		Description: "Queries the Wolfram|Alpha API about the input.",
		Level:       "100",
		Usage:       "()wa <input>",
	},
	"weather": {
		Name:        "weather",
		Alias:       nil,
		Description: `Returns the weather for a given location. If you "$set location" your location, the command will use that location if no other location is specified.`,
		Level:       "0",
		Usage:       "()weather [location]",
	},
	"xkcd": {
		Name:        "xkcd",
		Alias:       nil,
		Description: "Returns the link to the latest xkcd comic.",
		Level:       "0",
		Usage:       "()xkcd",
	},
	"yaf": {
		Name:        "yaf",
		Alias:       nil,
		Description: "Downloads the video of a given link with yt-dlp and then uploads the video to yaf.li.",
		Level:       "420",
		Usage:       "()yaf <link>",
	},
}

// Help checks if a help text for a given command exists and replies with it.
func (app *application) commandHelp(target, name, username string, message twitch.PrivateMessage) {
	// Check if the `helpText` map has an entry for `name`. If it does return it's value entry
	// and send that as a reply.
	i, ok := helpText[name]
	if !ok {
		// If it doesn't check the database for a command with that `name`. If there is one
		// reply with that commands `help` entry.
		c, err := app.GetCommandDescription(name, target, username)
		if err != nil {
			app.Log.Infow("commandHelp: no such command found",
				"err", err)
			return
		}

		app.Send(target, c, message)
		return
	}

	app.Send(target, i.Description, message)
}

// Help checks if a help text for a given command exists and replies with it.
func (app *application) GetAllHelpText() string {
	// The slice of timers is only used to log them at
	// the start so it looks a bit nicer.
	var cs []string

	// Iterate over all timers and then add them onto the scheduler.
	for i, v := range helpText {
		// idk why this works but it does so no touchy touchy.
		// https://github.com/robfig/cron/issues/420#issuecomment-940949195
		i, v := i, v
		_ = i
		var c string

		if v.Alias == nil {
			c = fmt.Sprintf("Name: %s\nDescription: %s\nLevel: %s\nUsage: %s\n\n", i, v.Description, v.Level, v.Usage)
		} else {
			c = fmt.Sprintf("Name: %s\nAliases: %s\nDescription: %s\nLevel: %s\nUsage: %s\n\n", i, v.Alias, v.Description, v.Level, v.Usage)

		}

		// Add new value to the slice
		cs = append(cs, c)
	}

	sort.Strings(cs)
	return strings.Join(cs, "")
}
