package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/lyx0/nourybot/internal/config"
	"github.com/lyx0/nourybot/pkg/models/mysql"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type nourybot struct {
	twitchClient *twitch.Client
	uptime       time.Time
	channels     *mysql.ChannelModel
}

func main() {
	cfg := config.LoadConfig()
	logrus.Info(cfg.Username, cfg.Oauth)

	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)
	now := time.Now()

	db, err := openDB(cfg.MySQLURI)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	nb := &nourybot{
		twitchClient: twitchClient,
		uptime:       now,
		channels:     &mysql.ChannelModel{DB: db},
	}

	nb.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		nb.onPrivateMessage(message)
	})

	nb.twitchClient.OnConnect(nb.onConnect)

	err = nb.connect()
	if err != nil {
		panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
