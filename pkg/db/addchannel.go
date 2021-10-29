package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
)

// AddChannel adds a given channel name to the database.
func AddChannel(target, channelName string, nb *bot.Bot) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := nb.MongoClient.Database("nourybot").Collection("channels")

	// Channel
	// {{"name": string}, {"connect": bool}}
	chnl := Channel{channelName, true}

	_, insertErr := collection.InsertOne(ctx, chnl)
	if insertErr != nil {
		nb.Send(target, fmt.Sprintf("Error trying to join %s", channelName))
		log.Error(insertErr)
		return
	}
	nb.TwitchClient.Join(channelName)
	nb.Send(target, fmt.Sprintf("Joined %s", channelName))
}
