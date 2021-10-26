package handlers

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/utils"
)

// Command contains all the logic for routing mesasges containing commands
// and will forward the messages to the specific command handlers.
func Command(message twitch.PrivateMessage, nb *bot.Bot) {
	// logrus.Info("fn Command")

	utils.CommandUsed()

	// commandName is the actual command name without the prefix.
	commandName := strings.ToLower(strings.SplitN(message.Message, " ", 3)[0][2:])

	// cmdParams are additional command inputs.
	// example: "weather san antonion
	// is: commandName cmdParams[0] cmdParams[1]
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
			return
		}
	case "7tv":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()7tv [emote]")
			return
		}
		commands.SevenTV(target, cmdParams[1], nb)
		return

	case "bot":
		commands.Help(target, nb)
		return

	case "botinfo":
		commands.Help(target, nb)
		return

	case "botstatus":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()botstatus [username]")
			return
		} else {
			commands.BotStatus(target, cmdParams[1], nb)
			return
		}

	case "bttv":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()bttv [emote]")
			return
		}
		commands.Bttv(target, cmdParams[1], nb)
		return

	// case "bttvemotes":
	// 	commands.BttvEmotes(target, nb)
	// 	return

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

	case "commands":
		commands.CommandsList(target, nb)
		return

	case "mycolor":
		commands.Color(message, nb)
		return

	case "currency":
		if msgLen <= 4 {
			nb.Send(target, "Usage: ()currency 10 usd to eur, only 3 letter codes work.")
			return
		}
		commands.Currency(target, cmdParams[1], cmdParams[2], cmdParams[4], nb)
		return

	case "echo":
		commands.Echo(target, message.Message[7:len(message.Message)], nb)
		return

	case "8ball":
		commands.EightBall(target, nb)
		return

	case "emote":
		commands.EmoteLookup(target, cmdParams[1], nb)

	case "emotelookup":
		commands.EmoteLookup(target, cmdParams[1], nb)

	case "ffz":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()ffz [emote]")
			return
		}
		commands.Ffz(target, cmdParams[1], nb)

	// case "ffzemotes":
	// 	commands.FfzEmotes(target, nb)
	// 	return

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
			commands.Firstline(target, target, cmdParams[1], nb)
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
			commands.Firstline(target, target, cmdParams[1], nb)
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

	case "godoc":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()godoc [term]")
			return
		} else {
			commands.Godocs(target, message.Message[8:len(message.Message)], nb)
			return
		}

	case "godocs":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()godoc [term]")
			return
		} else {
			commands.Godocs(target, message.Message[9:len(message.Message)], nb)
			return
		}

	case "help":
		commands.Help(target, nb)

	case "nourybot":
		commands.Help(target, nb)

	case "num":
		if msgLen == 1 {
			commands.RandomNumber(target, nb)
		} else {
			commands.Number(target, cmdParams[1], nb)
		}

	case "number":
		if msgLen == 1 {
			commands.RandomNumber(target, nb)
		} else {
			commands.Number(target, cmdParams[1], nb)
		}
	case "osrs":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()osrs [term]")
			return
		} else {
			commands.Osrs(target, message.Message[7:len(message.Message)], nb)
			return
		}

	case "ping":
		commands.Ping(target, nb)
		return

	case "pingme":
		commands.Pingme(target, message.User.DisplayName, nb)
		return

	case "preview":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()preview [channel]")
			return
		} else {
			commands.Thumbnail(target, cmdParams[1], nb)
			return
		}

	case "profilepicture":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()profilepicture [user]")
			return
		}

		commands.ProfilePicture(target, cmdParams[1], nb)
		return

	case "pfp":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()pfp [user]")
			return
		}

		commands.ProfilePicture(target, cmdParams[1], nb)
		return

	// case "pyramid":
	// 	if msgLen != 3 {
	// 		nb.Send(target, "Usage: ()pyramid [size] [emote]")
	// 	} else if utils.ElevatedPrivsMessage(message) {
	// 		commands.Pyramid(target, cmdParams[1], cmdParams[2], nb)
	// 	} else {
	// 		nb.Send(target, "Pleb's can't pyramid FeelsBadMan")
	// 	}

	case "randomcat":
		commands.RandomCat(target, nb)
		return

	case "cat":
		commands.RandomCat(target, nb)
		return

	case "randomdog":
		commands.RandomDog(target, nb)
		return

	case "dog":
		commands.RandomDog(target, nb)
		return

	case "randomduck":
		commands.RandomDuck(target, nb)
		return

	case "duck":
		commands.RandomDuck(target, nb)
		return

	case "fox":
		commands.RandomFox(target, nb)
		return

	case "randomfox":
		commands.RandomFox(target, nb)
		return

	case "rq":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()rq [channel] [user]")
			return
		} else if msgLen == 2 {
			commands.RandomQuote(target, target, cmdParams[1], nb)
			return
		} else {
			commands.RandomQuote(target, cmdParams[1], cmdParams[2], nb)
			return
		}

	case "randomquote":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()randomquote [channel] [user]")
			return
		} else if msgLen == 2 {
			commands.RandomQuote(target, target, cmdParams[1], nb)
			return
		} else {
			commands.RandomQuote(target, cmdParams[1], cmdParams[2], nb)
			return
		}

	case "randomxkcd":
		commands.RandomXkcd(target, nb)
		return

	case "robo":
		commands.RoboHash(target, message, nb)
		return

	case "robohash":
		commands.RoboHash(target, message, nb)
		return

	case "subage":
		if msgLen < 3 {
			nb.Send(target, "Usage: ()subage [streamer] [user]")
			return
		} else {
			commands.Subage(target, cmdParams[2], cmdParams[1], nb)
			return
		}

	case "thumb":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()thumb [channel]")
			return
		} else {
			commands.Thumbnail(target, cmdParams[1], nb)
			return
		}

	case "thumbnail":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()thumbnail [channel]")
			return
		} else {
			commands.Thumbnail(target, cmdParams[1], nb)
			return
		}

	case "title":
		if msgLen == 1 {
			commands.Title(target, target, nb)
			return
		} else {
			commands.Title(target, cmdParams[1], nb)
			return
		}

	case "uptime":
		if msgLen == 1 {
			commands.Uptime(target, target, nb)
			return
		} else {
			commands.Uptime(target, cmdParams[1], nb)
			return
		}

	case "uid":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()uid [username]")
			return
		} else {
			commands.Userid(target, cmdParams[1], nb)
			return
		}

	case "userid":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()userid [username]")
			return
		} else {
			commands.Userid(target, cmdParams[1], nb)
			return
		}

	case "weather":
		if msgLen == 1 {
			nb.Send(target, "Usage: ()weather [location]")
			return
		} else {
			commands.Weather(target, message.Message[9:len(message.Message)], nb)
			return
		}

	case "xd":
		commands.Xd(target, nb)
		return

	case "xkcd":
		commands.Xkcd(target, nb)
		return
	}
}
