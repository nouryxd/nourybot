package bot

import (
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

type Bot struct {
	config       config
	twitchClient *twitch.Client
	logger       *logrus.Logger
}

func New() *Bot {
	var cfg config

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

	bot := &Bot{
		config:       cfg,
		twitchClient: twitchClient,
		logger:       lgr,
	}
	// Received a PrivateMessage (normal chat message), pass it to
	// the handler who checks for further action.
	bot.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		bot.handlePrivateMessage(message)
	})

	// Received a WhisperMessage (Twitch DM), pass it to
	// the handler who checks for further action.
	bot.twitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		bot.handleWhisperMessage(message)
	})

	// Successfully connected to Twitch so we log a message with the
	// mode we are currently running in..
	bot.twitchClient.OnConnect(func() {
		bot.logger.Infof("Successfully connected to Twitch Servers in %s mode!", bot.config.env)
	})

	return bot

}
