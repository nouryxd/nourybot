package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lyx0/nourybot/internal/data"
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

func main() {
	var cfg config

	// Initialize a new sugared logger that we'll pass on
	// down through the application.
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Twitch
	cfg.twitchUsername = os.Getenv("TWITCH_USERNAME")
	cfg.twitchOauth = os.Getenv("TWITCH_OAUTH")
	tc := twitch.NewClient(cfg.twitchUsername, cfg.twitchOauth)

	// Environment
	cfg.environment = "Development"

	// Database
	cfg.db.dsn = os.Getenv("DB_DSN")
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = "15m"

	db, err := openDB(cfg)
	if err != nil {
		sugar.Fatal(err)
	}

	// Initialize Application
	app := &Application{
		TwitchClient: tc,
		Logger:       sugar,
		Db:           db,
		Models:       data.NewModels(db),
	}

	// Received a PrivateMessage (normal chat message).
	app.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		app.handlePrivateMessage(message)
	})

	// Received a WhisperMessage (Twitch DM).
	app.TwitchClient.OnWhisperMessage(func(message twitch.WhisperMessage) {
		app.handleWhisperMessage(message)
	})

	// Successfully connected to Twitch
	app.TwitchClient.OnConnect(func() {
		app.Logger.Infow("Successfully connected to Twitch Servers",
			"Bot username", cfg.twitchUsername,
			"Environment", cfg.environment,
			"Database Open Conns", cfg.db.maxOpenConns,
			"Database Idle Conns", cfg.db.maxIdleConns,
			"Database Idle Time", cfg.db.maxIdleTime,
			"Database", db.Stats(),
		)

		// Start time
		common.StartTime()

		common.Send("nourylul", "xd", app.TwitchClient)
		// app.GetAllChannels()
		app.InitialJoin()
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
