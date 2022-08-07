package main

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/commands"
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

func handleCommand(message twitch.PrivateMessage, tc *twitch.Client) {
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
			common.Send(target, "xd", tc)
			return
		}
	case "currency":
		if msgLen < 4 {
			common.Send(target, "Not enough arguments provided. Usage: ()currency 10 USD to EUR", tc)
			return
		}
		commands.Currency(target, cmdParams[1], cmdParams[2], cmdParams[4], tc)
	case "echo":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided.", tc)
			return
		} else {
			commands.Echo(target, message.Message[7:len(message.Message)], tc)
			return
		}
	case "tweet":
		if msgLen < 2 {
			common.Send(target, "Not enough arguments provided. Usage: ()tweet forsen", tc)
			return
		} else {
			commands.Tweet(target, cmdParams[1], tc)
		}
	case "ping":
		commands.Ping(target, tc)
	}
}
