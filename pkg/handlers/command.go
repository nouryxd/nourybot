package handlers

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/sirupsen/logrus"
)

func Command(message twitch.PrivateMessage, nb *bot.Bot) {
	logrus.Info("fn Command")

	// commandName is the actual command name without the prefix.
	commandName := strings.ToLower(strings.SplitN(message.Message, " ", 3)[0][2:])

	// cmdParams are additional command inputs.
	// example:
	// weather san antonio
	// is
	// commandName cmdParams[0] cmdParams[1]
	cmdParams := strings.SplitN(message.Message, " ", 500)

	// msgLen is the amount of words in the message without the prefix.
	// Useful for checking if enough cmdParams are given.
	msgLen := len(strings.SplitN(message.Message, " ", -2))

	// target channel
	target := message.Channel

	switch commandName {
	case "":
		if msgLen == 1 {
			nb.Send(target, "xd")
		}
	case "botstatus":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()botstatus [username]")
			return
		} else {
			commands.BotStatus(target, cmdParams[1], nb)
			return
		}
	case "bttvemotes":
		commands.BttvEmotes(target, nb)
		return
	case "cf":
		commands.Coinflip(target, nb)
		return
	case "coin":
		commands.Coinflip(target, nb)
		return
	case "coinflip":
		commands.Coinflip(target, nb)
		return
	case "color":
		commands.Color(message, nb)
		return
	case "mycolor":
		commands.Color(message, nb)
		return
	case "echo":
		commands.Echo(target, message.Message[7:len(message.Message)], nb)
		return
	case "8ball":
		commands.EightBall(target, nb)
		return
	case "ffzemotes":
		commands.FfzEmotes(target, nb)
		return

	case "fill":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()fill [emote]")
			return
		} else {
			commands.Fill(target, message.Message[7:len(message.Message)], nb)
			return
		}
	case "firstline":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()firstline [channel] [user]")
			return
		} else if msgLen == 2 {
			commands.Firstline(target, message.Channel, cmdParams[1], nb)
			return
		} else {
			commands.Firstline(target, cmdParams[1], cmdParams[2], nb)
			return
		}
	case "fl":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()firstline [channel] [user]")
			return
		} else if msgLen == 2 {
			commands.Firstline(target, message.Channel, cmdParams[1], nb)
			return
		} else {
			commands.Firstline(target, cmdParams[1], cmdParams[2], nb)
			return
		}
	case "followage":
		if msgLen <= 2 {
			nb.Send(target, "Usage: ()followage [channel] [user]")
			return
		} else {
			commands.Followage(target, cmdParams[1], cmdParams[2], nb)
			return
		}
	case "game":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()game [channel]")
			return
		} else {
			commands.Game(target, cmdParams[1], nb)
		}

	case "randomxkcd":
		commands.RandomXkcd(target, nb)
		return
	case "ping":
		commands.Ping(target, nb)
		return

	case "weather":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()weather [location]")
			return
		} else {
			commands.Weather(target, message.Message[9:len(message.Message)], nb)
			return
		}
	case "xkcd":
		commands.Xkcd(target, nb)
		return
	}

}
