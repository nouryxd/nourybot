package db

import (
	"context"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/sirupsen/logrus"
)

type logMessage struct {
	LoginName string `json:"login_name"`
	Channel   string `json:"channel"`
	UserID    string `json:"user_id"`
	Text      string `json:"message"`
}

func (l *logMessage) insert(nb *bot.Bot, name, channel, id, text string) error {
	logrus.Info("logging message")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	collection := nb.MongoClient.Database("nourybot").Collection("messageLogs")

	_, err := collection.InsertOne(ctx, &logMessage{
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

// InsertMessage is called on every message and logs message details
// into the database.
func InsertMessage(nb *bot.Bot, name, channel, id, text string) {
	err := (&logMessage{}).insert(nb, name, channel, id, text)
	if err != nil {
		logrus.Info("failed to log message")
	}
}
