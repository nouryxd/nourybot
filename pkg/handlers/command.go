package handlers

import (
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// HandleCommand receives a twitch.PrivateMessage from
// HandlePrivateMessage where it found a command in it.
// HandleCommand passes on the message to the specific
// command handler for further action.
func HandleCommand(message twitch.PrivateMessage, twitchClient *twitch.Client, uptime time.Time) {
	log.Info("fn HandleCommand")

	// Counter that increments on every command call.
	utils.CommandUsed()

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

	switch commandName {
	case "":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Why yes, that's my prefix :)")
			return
		}
	case "botstatus":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()botstatus [username]")
			return
		} else {
			commands.BotStatus(message.Channel, cmdParams[1], twitchClient)
			return
		}
	case "bttvemotes":
		commands.BttvEmotes(message.Channel, twitchClient)
		return
	case "cf":
		commands.Coinflip(message.Channel, twitchClient)
		return
	case "coin":
		commands.Coinflip(message.Channel, twitchClient)
		return
	case "coinflip":
		commands.Coinflip(message.Channel, twitchClient)
		return
	case "color":
		commands.Color(message, twitchClient)
		return
	case "mycolor":
		commands.Color(message, twitchClient)
		return
	case "echo":
		commands.Echo(message.Channel, message.Message[7:len(message.Message)], twitchClient)
		return
	case "8ball":
		commands.EightBall(message.Channel, twitchClient)
		return
	case "ffzemotes":
		commands.FfzEmotes(message.Channel, twitchClient)
		return
	case "fill":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()fill [emote]")
			return
		} else {
			commands.Fill(message.Channel, message.Message[7:len(message.Message)], twitchClient)
			return
		}
	case "firstline":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()firstline [channel] [user]")
			return
		} else if msgLen == 2 {
			commands.Firstline(message.Channel, message.Channel, cmdParams[1], twitchClient)
			return
		} else {
			commands.Firstline(message.Channel, cmdParams[1], cmdParams[2], twitchClient)
			return
		}
	case "fl":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()firstline [channel] [user]")
			return
		} else if msgLen == 2 {
			commands.Firstline(message.Channel, message.Channel, cmdParams[1], twitchClient)
			return
		} else {
			commands.Firstline(message.Channel, cmdParams[1], cmdParams[2], twitchClient)
			return
		}
	case "followage":
		if msgLen <= 2 {
			twitchClient.Say(message.Channel, "Usage: ()followage [channel] [user]")
			return
		} else {
			commands.Followage(message.Channel, cmdParams[1], cmdParams[2], twitchClient)
			return
		}
	case "game":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()game [channel]")
			return
		} else {
			commands.Game(message.Channel, cmdParams[1], twitchClient)
		}
	case "godoc":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()godoc [term]")
			return
		} else {
			commands.Godocs(message.Channel, message.Message[8:len(message.Message)], twitchClient)
			return
		}
	case "godocs":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()godoc [term]")
			return
		} else {
			commands.Godocs(message.Channel, message.Message[9:len(message.Message)], twitchClient)
			return
		}
	case "num":
		if msgLen == 1 {
			commands.RandomNumber(message.Channel, twitchClient)
		} else {
			commands.Number(message.Channel, cmdParams[1], twitchClient)
		}
	case "number":
		if msgLen == 1 {
			commands.RandomNumber(message.Channel, twitchClient)
		} else {
			commands.Number(message.Channel, cmdParams[1], twitchClient)
		}
	case "ping":
		commands.Ping(message.Channel, twitchClient, uptime)
		return
	case "pingme":
		commands.Pingme(message.Channel, message.User.DisplayName, twitchClient)
		return
	case "pyramid":
		if msgLen != 3 {
			twitchClient.Say(message.Channel, "Usage: ()pyramid [size] [emote]")
		} else if utils.ElevatedPrivsMessage(message) {
			commands.Pyramid(message.Channel, cmdParams[1], cmdParams[2], twitchClient)
		} else {
			twitchClient.Say(message.Channel, "Pleb's can't pyramid FeelsBadMan")
		}
	case "randomcat":
		commands.RandomCat(message.Channel, twitchClient)
		return
	case "randomfox":
		commands.RandomFox(message.Channel, twitchClient)
		return
	case "subage":
		if msgLen < 3 {
			twitchClient.Say(message.Channel, "Usage: ()subage [user] [streamer]")
			return
		} else {
			commands.Subage(message.Channel, cmdParams[1], cmdParams[2], twitchClient)
			return
		}
	case "thumb":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()thumbnail [channel]")
			return
		} else {
			commands.Thumbnail(message.Channel, cmdParams[1], twitchClient)
			return
		}
	case "thumbnail":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()thumbnail [channel]")
			return
		} else {
			commands.Thumbnail(message.Channel, cmdParams[1], twitchClient)
			return
		}
	case "uptime":
		if msgLen == 1 {
			commands.Uptime(message.Channel, message.Channel, twitchClient)
			return
		} else {
			commands.Uptime(message.Channel, cmdParams[1], twitchClient)
			return
		}
	case "weather":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()weather [location]")
			return
		} else {
			commands.Weather(message.Channel, message.Message[9:len(message.Message)], twitchClient)
			return
		}
	case "xd":
		commands.Xd(message.Channel, twitchClient)
		return

	}

}
