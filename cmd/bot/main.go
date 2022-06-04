package main

import (
	"flag"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type config struct {
	env         string
	botUsername string
	botOauth    string
}

type application struct {
	config       config
	twitchClient *twitch.Client
	logger       *logrus.Logger
}

func main() {
	var cfg config

	// Parse which environment we are running in. This will decide in
	// the future how many channels we join or which database we are
	// connecting to for example.
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|string|production)")
	flag.Parse()

	// Initialize a new logger we attach to our application struct.
	lgr := logrus.New()

	// Load the .env file and check for errors.
	err := godotenv.Load()
	if err != nil {
		lgr.Fatal("Error loading .env file")
	}

	// Load bot credentials from the .env file.
	cfg.botUsername = os.Getenv("BOT_USER")
	cfg.botOauth = os.Getenv("BOT_OAUTH")

	// Initialize a new twitch client which we attach to our
	// application struct.
	twitchClient := twitch.NewClient(cfg.botUsername, cfg.botOauth)

	// Finally Initialize a new application instance with our
	// attached methods.
	app := &application{
		config:       cfg,
		twitchClient: twitchClient,
		logger:       lgr,
	}

	// Received a PrivateMessage (normal chat message), pass it to
	// the handler who checks for further action.
	app.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		app.handlePrivateMessage(message)
	})

	// Received a WhisperMessage (Twitch DM), pass it to
	// the handler who checks for further action.
	app.twitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		app.handleWhisperMessage(message)
	})

	// Successfully connected to Twitch so we log a message with the
	// mode we are currently running in..
	app.twitchClient.OnConnect(func() {
		app.logger.Infof("Successfully connected to Twitch Servers in %s mode!", app.config.env)
	})

	// Join test channels
	app.twitchClient.Join("nourylul")
	app.twitchClient.Join("nourybot")

	// Say hello because we are nice :^)
	app.twitchClient.Say("nourylul", "RaccAttack")
	app.twitchClient.Say("nourybot", "RaccAttack")

	// Connect to the twitch IRC servers.
	err = app.twitchClient.Connect()
	if err != nil {
		panic(err)
	}
}
