package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/jakecoffman/cron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nicklaw5/helix/v2"
	"github.com/nouryxd/nourybot/internal/data"
	"github.com/nouryxd/nourybot/pkg/common"

	"go.uber.org/zap"
)

type config struct {
	twitchUsername     string
	twitchOauth        string
	twitchClientId     string
	twitchClientSecret string
	eventSubSecret     string
	twitchID           string
	wolframAlphaAppID  string
	commandPrefix      string
	env                string
	db                 struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	TwitchClient *twitch.Client
	HelixClient  *helix.Client
	Log          *zap.SugaredLogger
	Db           *sql.DB
	Models       data.Models
	Scheduler    *cron.Cron
	Environment  string
	Config       config
	// Rdb       *redis.Client
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	var cfg config

	logger := zap.NewExample()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Sugar().Errorw("error syncing logger",
				"error", err,
			)
		}
	}()
	sugar := logger.Sugar()

	common.StartTime()

	err := godotenv.Load()
	if err != nil {
		sugar.Fatal("Error loading .env")
	}

	// Twitch config variables
	cfg.twitchUsername = os.Getenv("TWITCH_USERNAME")
	cfg.twitchOauth = os.Getenv("TWITCH_OAUTH")
	cfg.twitchClientId = os.Getenv("TWITCH_CLIENT_ID")
	cfg.twitchClientSecret = os.Getenv("TWITCH_CLIENT_SECRET")
	cfg.eventSubSecret = os.Getenv("EVENTSUB_SECRET")
	cfg.wolframAlphaAppID = os.Getenv("WOLFRAMALPHA_APP_ID")
	cfg.twitchID = os.Getenv("TWITCH_ID")
	cfg.env = os.Getenv("ENV")
	tc := twitch.NewClient(cfg.twitchUsername, cfg.twitchOauth)

	switch cfg.env {
	case "dev":
		cfg.db.dsn = os.Getenv("DEV_DSN")
		cfg.commandPrefix = "!!"
	case "prod":
		cfg.db.dsn = os.Getenv("PROD_DSN")
		cfg.commandPrefix = "()"
	}
	// Database config variables
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = "15m"

	// Initialize a new Helix Client, request a new AppAccessToken
	// and set the token on the client.
	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:     cfg.twitchClientId,
		ClientSecret: cfg.twitchClientSecret,
	})
	if err != nil {
		sugar.Fatalw("Error creating new helix client",
			"helixClient", helixClient,
			"err", err,
		)
	}

	helixResp, err := helixClient.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		sugar.Fatalw("Helix: Error getting new helix AppAcessToken",
			"resp", helixResp,
			"err", err,
		)
	}

	// Set the access token on the client
	helixClient.SetAppAccessToken(helixResp.Data.AccessToken)

	// Establish database connection
	db, err := openDB(cfg)
	if err != nil {
		sugar.Fatalw("could not establish database connection",
			"err", err,
		)
	}
	app := &application{
		TwitchClient: tc,
		HelixClient:  helixClient,
		Log:          sugar,
		Db:           db,
		Models:       data.NewModels(db),
		Scheduler:    cron.New(),
		Config:       cfg,
		Environment:  cfg.env,
	}

	app.Log.Infow("db.Stats",
		"db.Stats", db.Stats(),
	)

	// Received a PrivateMessage (normal chat message).
	app.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// roomId is the Twitch UserID of the channel the message originated from.
		// If there is no roomId something went really wrong.
		roomId := message.Tags["room-id"]
		if roomId == "" {
			return
		}

		// Message was shorter than our prefix is therefore it's irrelevant for us.
		if len(message.Message) >= 2 {
			// This bots prefix is "()" configured above at cfg.commandPrefix,
			// Check if the first 2 characters of the mesage were our prefix.
			// if they were forward the message to the command handler.
			if message.Message[:2] == cfg.commandPrefix {
				go app.handleCommand(message)
				return
			}

			// Special rule for #pajlada.
			if message.Message == "!nourybot" {
				app.Send(message.Channel, "Lidl Twitch bot made by @nouryxd. Prefix: ()", message)
			}
		}
	})

	app.TwitchClient.OnClearChatMessage(func(message twitch.ClearChatMessage) {
		if message.BanDuration == 0 && message.Channel == "forsen" {
			app.TwitchClient.Say("nouryxd",
				fmt.Sprintf("MODS https://logs.ivr.fi/?channel=forsen&username=%v",
					message.TargetUsername))
		}

		if message.BanDuration >= 28700 && message.Channel == "forsen" {
			app.TwitchClient.Say("nouryxd",
				fmt.Sprintf("monkaS -%v https://logs.ivr.fi/?channel=forsen&username=%v",
					message.BanDuration, message.TargetUsername))
		}

	})

	app.InitialJoin()
	// Load the initial timers from the database.
	app.InitialTimers()

	// Start the timers.
	app.Scheduler.Start()

	app.TwitchClient.OnConnect(func() {
		app.TwitchClient.Say("nouryxd", "gopherDance")
		app.TwitchClient.Say("nourybot", "gopherDance")
		app.TwitchClient.Say("xnoury", "alienPls")
		app.TwitchClient.Say("uudelleenkykeytynyt", "pajaDink")

		// Successfully connected to Twitch
		app.Log.Infow("Successfully connected to Twitch Servers",
			"Bot username", cfg.twitchUsername,
			"Environment", cfg.env,
			"Database Open Conns", cfg.db.maxOpenConns,
			"Database Idle Conns", cfg.db.maxIdleConns,
			"Database Idle Time", cfg.db.maxIdleTime,
			"Database", db.Stats(),
			"Helix", helixResp,
		)

	})

	// Start status page
	go app.startRouter()

	// Actually connect to chat.
	err = app.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}

	return nil
}

// openDB returns the sql.DB connection pool.
func openDB(cfg config) (*sql.DB, error) {
	// sql.Open() creates an empty connection pool with the provided DSN
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set database restraints.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Parse the maxIdleTime string into an actual duration and set it.
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// Create a new context with a 5 second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// db.PingContext() is needed to actually check if the
	// connection to the database was successful.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil

}
