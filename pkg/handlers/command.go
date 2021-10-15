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
		}
		return
	case "xd":
		twitchClient.Say(message.Channel, "xd")
		return
	case "echo":
		commands.Echo(message.Channel, message.Message[7:len(message.Message)], twitchClient)
		return
	case "ping":
		commands.Ping(message.Channel, twitchClient, uptime)
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
	case "8ball":
		commands.EightBall(message, twitchClient)
		return
	case "weather":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()weather [location]")
			return
		} else {
			commands.Weather(message.Channel, message.Message[9:len(message.Message)], twitchClient)
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

	case "fill":
		if msgLen == 1 {
			twitchClient.Say(message.Channel, "Usage: ()fill [emote]")
			return
		} else {
			commands.Fill(message.Channel, message.Message[7:len(message.Message)], twitchClient)
			return
		}

	case "pyramid":
		if msgLen != 3 {
			twitchClient.Say(message.Channel, "Usage: ()pyramid [size] [emote]")
		} else if utils.ElevatedPrivsMessage(message) {
			commands.Pyramid(message.Channel, cmdParams[1], cmdParams[2], twitchClient)
		} else {
			twitchClient.Say(message.Channel, "Pleb's can't pyramid FeelsBadMan")
		}

	case "pingme":
		commands.Pingme(message.Channel, message.User.DisplayName, twitchClient)

	}
}
