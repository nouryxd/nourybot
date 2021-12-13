package db

import (
	"context"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	"github.com/sirupsen/logrus"
)

type logCommand struct {
	CommandName string `json:"command_name"`
	LoginName   string `json:"login_name"`
	Channel     string `json:"channel"`
	UserID      string `json:"user_id"`
	Text        string `json:"message"`
}

func (l *logCommand) insert(nb *bot.Bot, commandName, name, channel, id, text string) error {
	logrus.Info("logging command")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	collection := nb.MongoClient.Database("nourybot").Collection("commandLogs")

	_, err := collection.InsertOne(ctx, &logCommand{
		LoginName:   name,
		Channel:     channel,
		UserID:      id,
		Text:        text,
		CommandName: commandName,
	})
	if err != nil {
		logrus.Info("failed to log command")
		return err
	}
	return nil
}

// InsertCommand is called on every message that invokes a command
// and inserts the message and command details into the database.
func InsertCommand(nb *bot.Bot, commandName, name, channel, id, text string) {
	err := (&logCommand{}).insert(nb, commandName, name, channel, id, text)
	if err != nil {
		logrus.Info("failed to log command")
	}
}
