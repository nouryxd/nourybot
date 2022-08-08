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
	"github.com/lyx0/nourybot/pkg/common"
	"go.uber.org/zap"
)

type config struct {
	twitchUsername string
	twitchOauth    string
	environment    string
	db             struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type Application struct {
	TwitchClient *twitch.Client
	Logger       *zap.SugaredLogger
	Db           *sql.DB
}

func main() {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.twitchUsername = os.Getenv("TWITCH_USERNAME")
	cfg.twitchOauth = os.Getenv("TWITCH_OAUTH")
	cfg.environment = "Development"
	cfg.db.dsn = os.Getenv("DB_DSN")
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = "15m"

	tc := twitch.NewClient(cfg.twitchUsername, cfg.twitchOauth)
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	// Start time
	common.StartTime()

	db, err := openDB(cfg)
	if err != nil {
		sugar.Fatal(err)
	}
	app := &Application{
		TwitchClient: tc,
		Logger:       sugar,
		Db:           db,
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
		)

		common.Send("nourylul", "xd", app.TwitchClient)
	})
	app.TwitchClient.Join("nourylul")

	err = app.TwitchClient.Connect()
	if err != nil {
		panic(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
