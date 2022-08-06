package bot

import (
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	TwitchClient *twitch.Client
	logger       *logrus.Logger
}

func New() *Bot {
	// Initialize a new logger we attach to our application struct.
	lgr := logrus.New()

	cfg := config.New()

	// Initialize a new twitch client which we attach to our
	// application struct.
	logrus.Info(cfg.TwitchUsername, cfg.TwitchOauth)
	twitchClient := twitch.NewClient(cfg.TwitchUsername, cfg.TwitchOauth)

	bot := &Bot{
		TwitchClient: twitchClient,
		logger:       lgr,
	}

	// Received a PrivateMessage (normal chat message), pass it to
	// the handler who checks for further action.
	bot.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		bot.handlePrivateMessage(message)
	})

	// Received a WhisperMessage (Twitch DM), pass it to
	// the handler who checks for further action.
	bot.TwitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		bot.handleWhisperMessage(message)
	})

	// Successfully connected to Twitch so we log a message with the
	// mode we are currently running in..
	bot.TwitchClient.OnConnect(func() {
		bot.logger.Infof("Successfully connected to Twitch Servers in %s mode!", cfg.Environment)
		bot.Send("nourylul", "xd")
	})

	return bot

}
