package main

import (
	"flag"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
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

	cfg.botUsername = os.Getenv("TWITCH_USER")
	cfg.botOauth = os.Getenv("TWITCH_OAUTH")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	twitchClient := twitch.NewClient(cfg.botUsername, cfg.botOauth)
	app := &application{
		config:       cfg,
		twitchClient: twitchClient,
		logger:       logger,
	}

	app.twitchClient.Connect()
}
