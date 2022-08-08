package main

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

func handleCommand(message twitch.PrivateMessage, app *Application) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	// Increments the counter how many commands were used.
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
	_ = cmdParams

	// msgLen is the amount of words in a message without the prefix.
	// Useful to check if enough cmdParams are provided.
	msgLen := len(strings.SplitN(message.Message, " ", -2))

	// target is the channelname the message originated from and
	// where we are responding.
	target := message.Channel

	sugar.Infow("Command received",
		// "message", message,
		"message.Channel", target,
		"message.Message", message.Message,
		"commandName", commandName,
		"cmdParams", cmdParams,
		"msgLen", msgLen,
	)

	switch commandName {
	case "":
		if msgLen == 1 {
			common.Send(target, "xd", app.TwitchClient)
			return
		}
	case "addchannel":
		if message.User.ID != "31437432" { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			// ()addchannel noemience
			AddChannel(cmdParams[1], message, app)
			return
		}
	case "bttv":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()bttv <emote name>", app.TwitchClient)
			return
		} else {
			commands.Bttv(target, cmdParams[1], app.TwitchClient)
			return
		}
	case "bttvemotes": // Pretty spammy command so it's limited to my own channel until elevated messages are implemented.
		if target == "nourylul" || target == "nourybot" {
			commands.Bttvemotes(target, app.TwitchClient)
			return
		} else {
			return
		}
	case "coin":
		commands.Coinflip(target, app.TwitchClient)
		return
	case "coinflip":
		commands.Coinflip(target, app.TwitchClient)
		return
	case "cf":
		commands.Coinflip(target, app.TwitchClient)
		return
	case "currency":
		if msgLen < 4 {
			common.Send(target, "Not enough arguments provided. Usage: ()currency 10 USD to EUR", app.TwitchClient)
			return
		}
		commands.Currency(target, cmdParams[1], cmdParams[2], cmdParams[4], app.TwitchClient)
		return
	case "echo":
		if message.User.ID != "31437432" { // Limit to myself for now.
			return
		} else if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", app.TwitchClient)
			return
		} else {
			commands.Echo(target, message.Message[7:len(message.Message)], app.TwitchClient)
			return
		}

	case "ffz":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()ffz <emote name>", app.TwitchClient)
			return
		} else {
			commands.Ffz(target, cmdParams[1], app.TwitchClient)
			return
		}
	case "ffzemotes": // Pretty spammy command so it's limited to my own channel until elevated messages are implemented.
		if target == "nourylul" || target == "nourybot" {
			commands.Ffzemotes(target, app.TwitchClient)
			return
		} else {
			return
		}

		// Followage
		// ()followage channel username
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
	case "preview":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()preview <username>", app.TwitchClient)
			return
		} else {
			commands.Preview(target, cmdParams[1], app.TwitchClient)
			return
		}

	case "thumbnail":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()thumbnail <username>", app.TwitchClient)
			return
		} else {
			commands.Preview(target, cmdParams[1], app.TwitchClient)
			return
		}

	// SevenTV
	case "seventv":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()seventv <emote name>", app.TwitchClient)
			return
		} else {
			commands.Seventv(target, cmdParams[1], app.TwitchClient)
			return
		}
	case "7tv":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()seventv <emote name>", app.TwitchClient)
			return
		} else {
			commands.Seventv(target, cmdParams[1], app.TwitchClient)
			return
		}
	case "tweet":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()tweet <username>", app.TwitchClient)
			return
		} else {
			commands.Tweet(target, cmdParams[1], app.TwitchClient)
		}
	case "rxkcd":
		commands.RandomXkcd(target, app.TwitchClient)
		return
	case "randomxkcd":
		commands.RandomXkcd(target, app.TwitchClient)
		return
	case "xkcd":
		commands.Xkcd(target, app.TwitchClient)
		return
	}
}
