package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/jakecoffman/cron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lyx0/nourybot/internal/common"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/nicklaw5/helix"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type config struct {
	twitchUsername     string
	twitchOauth        string
	twitchClientId     string
	twitchClientSecret string
	commandPrefix      string
	db                 struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type Application struct {
	TwitchClient *twitch.Client
	HelixClient  *helix.Client
	Logger       *zap.SugaredLogger
	Db           *sql.DB
	Models       data.Models
	Scheduler    *cron.Cron
	Rdb          *redis.Client
}

var envFlag string
var ctx = context.Background()

func init() {
	flag.StringVar(&envFlag, "env", "dev", "database connection to use: (dev/prod)")
	flag.Parse()
}

func main() {
	var cfg config

	// Initialize a new sugared logger that we'll pass on
	// down through the application.
	logger := zap.NewExample()
	defer logger.Sync()
	sugar := logger.Sugar()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Twitch config variables
	cfg.twitchUsername = os.Getenv("TWITCH_USERNAME")
	cfg.twitchOauth = os.Getenv("TWITCH_OAUTH")
	cfg.twitchClientId = os.Getenv("TWITCH_CLIENT_ID")
	cfg.twitchClientSecret = os.Getenv("TWITCH_CLIENT_SECRET")
	cfg.commandPrefix = os.Getenv("TWITCH_COMMAND_PREFIX")
	tc := twitch.NewClient(cfg.twitchUsername, cfg.twitchOauth)

	switch envFlag {
	case "dev":
		cfg.db.dsn = os.Getenv("LOCAL_DSN")
	case "prod":
		cfg.db.dsn = os.Getenv("SUPABASE_DSN")
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
	sugar.Infow("Got new helix AppAccessToken",
		"helixClient", helixResp,
	)

	// Set the access token on the client
	helixClient.SetAppAccessToken(helixResp.Data.AccessToken)

	// Establish database connection
	db, err := openDB(cfg)
	if err != nil {
		sugar.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		sugar.Panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		sugar.Panic(err)
	}
	sugar.Infow("Redis initialization key",
		"key", val,
	)

	// Initialize Application with the new values
	app := &Application{
		TwitchClient: tc,
		HelixClient:  helixClient,
		Logger:       sugar,
		Db:           db,
		Models:       data.NewModels(db),
		Scheduler:    cron.New(),
		Rdb:          rdb,
	}

	// Received a PrivateMessage (normal chat message).
	app.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {

		// app.Logger.Infow("Message received",
		// 	"message", message,
		// 	"message.User.DisplayName", message.User.DisplayName,
		// 	"message.Message", message.Message,
		// )

		// roomId is the Twitch UserID of the channel the message originated from.
		// If there is no roomId something went really wrong.
		roomId := message.Tags["room-id"]
		if roomId == "" {
			app.Logger.Errorw("Missing room-id in message tag",
				"roomId", roomId,
			)
			return
		}

		// Message was shorter than our prefix is therefore it's irrelevant for us.
		if len(message.Message) >= 2 {
			// This bots prefix is "()" configured above at cfg.commandPrefix,
			// Check if the first 2 characters of the mesage were our prefix.
			// if they were forward the message to the command handler.
			if message.Message[:2] == cfg.commandPrefix {
				app.InitUser(message.User.Name, message.User.ID, message)
				app.handleCommand(message)
				return
			}

			// Special rule for #pajlada.
			if message.Message == "!nourybot" {
				common.Send(message.Channel, "Lidl Twitch bot made by @nourylul. Prefix: ()", app.TwitchClient)
			}

		}
	})

	// Received a WhisperMessage (Twitch DM).
	app.TwitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		// Print the whisper message for now.
		app.Logger.Infow("Whisper Message received",
			"message", message,
			"message.User.DisplayName", message.User.DisplayName,
			"message.Message", message.Message,
		)
	})

	// Successfully connected to Twitch
	app.TwitchClient.OnConnect(func() {
		app.Logger.Infow("Successfully connected to Twitch Servers",
			"Bot username", cfg.twitchUsername,
			"Environment", envFlag,
			"Database Open Conns", cfg.db.maxOpenConns,
			"Database Idle Conns", cfg.db.maxIdleConns,
			"Database Idle Time", cfg.db.maxIdleTime,
			"Database", db.Stats(),
			"Helix", helixResp,
		)

		// Start time
		common.StartTime()

		app.loadCommandHelp()

		// Join the channels in the database.
		app.InitialJoin()

		// Load the initial timers from the database.
		app.InitialTimers()

		// Start the timers.
		app.Scheduler.Start()

		common.Send("nourylul", "dankCircle", app.TwitchClient)
		common.Send("nourybot", "gopherDance", app.TwitchClient)
		common.Send("xnoury", "pajaDink", app.TwitchClient)
		common.Send("uudelleenkytkeytynyt", "PepeS", app.TwitchClient)
	})

	// Actually connect to chat.
	err = app.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}
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
