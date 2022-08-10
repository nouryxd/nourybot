package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func (app *Application) AddCommand(name string, message twitch.PrivateMessage) {
	// prefixLength is the length of `()addcommand` plus +2 (for the space and zero based)
	prefixLength := 14

	// Split the twitch message at the length of the prefix + the length of the name of the command.
	//      prefixLength |name| text
	//      0123456789012|4567|
	// e.g. ()addcommand dank FeelsDankMan
	//      |   part1    snip ^  part2   |
	text := message.Message[prefixLength+len(name) : len(message.Message)]
	err := app.Models.Commands.Insert(name, text)

	//	app.Logger.Infow("Message splits",
	//		"Command Name:", name,
	//		"Command Text:", text)

	if err != nil {
		reply := fmt.Sprintf("Something went wrong FeelsBadMan %s", err)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	} else {

		reply := fmt.Sprintf("Successfully added command: %s", name)
		common.Send(message.Channel, reply, app.TwitchClient)
		return
	}
}

func (app *Application) GetCommand(name string) (string, error) {
	command, err := app.Models.Commands.Get(name)

	if err != nil {
		return "", err
	}

	return command.Text, nil

}
