package main

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/sirupsen/logrus"
)

func (nb *nourybot) connect() error {
	nb.twitchClient.Join("nourybot")

	logrus.Info("Connecting to Twitch...")
	err := nb.twitchClient.Connect()
	if err != nil {
		panic(err)
	}

	return nil
}

func (nb *nourybot) onConnect() {
	nb.twitchClient.Say("nourybot", "xd")
}

func (nb *nourybot) onPrivateMessage(message twitch.PrivateMessage) {
	logrus.Info(message.Message)
}
