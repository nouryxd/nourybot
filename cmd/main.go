package main

import (
	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/lyx0/nourybot/pkg/config"
)

func main() {
	cfg := config.LoadConfig()
	nb := bot.NewBot(cfg)

	nb.Join("nourybot")

	nb.Say("nourybot", "HeyGuys")

	nb.Connect()
}
