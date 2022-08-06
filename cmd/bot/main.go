package main

import "github.com/lyx0/nourybot/pkg/bot"

func main() {

	bot := bot.New()

	bot.TwitchClient.Join("nourylul")
	err := bot.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}
}
