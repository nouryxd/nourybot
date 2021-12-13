package db

import (
	"context"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/sirupsen/logrus"
)

type LogMessage struct {
	LoginName string `json:"login_name"`
	Channel   string `json:"channel"`
	UserID    string `json:"user_id"`
	Text      string `json:"message"`
}

func (l *LogMessage) Insert(nb *bot.Bot, name, channel, id, text string) error {
	logrus.Info("logging message")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	collection := nb.MongoClient.Database("nourybot").Collection("logs")

	_, err := collection.InsertOne(ctx, &LogMessage{
		LoginName: name,
		Channel:   channel,
		UserID:    id,
		Text:      text,
	})
	if err != nil {
		logrus.Info("failed to log message")
		return err
	}
	return nil
}

func InsertMessage(nb *bot.Bot, name, channel, id, text string) {
	err := (&LogMessage{}).Insert(nb, name, channel, id, text)
	if err != nil {
		logrus.Info("failed to log message")
	}
}
