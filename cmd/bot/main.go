package main

import (
	"flag"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

type config struct {
	env         string
	botUsername string
	botOauth    string
}

type application struct {
	config       config
	twitchClient *twitch.Client
	logger       *log.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "development", "Environment (development|string|production)")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.botUsername = os.Getenv("BOT_USER")
	cfg.botOauth = os.Getenv("BOT_OAUTH")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	twitchClient := twitch.NewClient(cfg.botUsername, cfg.botOauth)
	app := &application{
		config:       cfg,
		twitchClient: twitchClient,
		logger:       logger,
	}

	app.twitchClient.Join("nourylul")
	app.twitchClient.Join("nourybot")

	app.twitchClient.Say("nourylul", "xd")
	app.twitchClient.Say("nourybot", "xd")

	err = app.twitchClient.Connect()
	if err != nil {
		panic(err)
	}
}
