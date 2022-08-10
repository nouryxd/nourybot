package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/pkg/common"
)

func (app *Application) AddCommand(name string, message twitch.PrivateMessage) {
	// text1 := strings.Split(message.Message, name)

	text := message.Message[14+len(name) : len(message.Message)]
	app.Logger.Infow("Message splits",
		"Command Name:", name,
		"Command Text:", text)
	err := app.Models.Commands.Insert(name, text)

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
