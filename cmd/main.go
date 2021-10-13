package main

import (
	"github.com/lyx0/nourybot/bot"
	"github.com/lyx0/nourybot/config"
)

func main() {
	cfg := config.LoadConfig()
	nb := bot.NewBot(cfg)

	nb.Join("nourybot")
	nb.Say("nourybot", "test")
	nb.Connect()
}
