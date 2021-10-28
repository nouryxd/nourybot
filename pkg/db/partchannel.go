package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
)

func PartChannel(target, channelName string, nb *bot.Bot) {
	// Interact with data

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/*
		Get my collection instance
	*/
	collection := nb.MongoClient.Database("nourybot").Collection("channels")

	// Channel
	chnl := Channel{channelName, true}

	res, insertErr := collection.DeleteOne(ctx, chnl)
	if insertErr != nil || res.DeletedCount != 1 {
		nb.Send(target, fmt.Sprintf("Error trying to part %s", channelName))
		log.Error(insertErr)
		return
	}
	nb.Send(target, fmt.Sprintf("Parted %s", channelName))

	log.Info(res.DeletedCount)

}
