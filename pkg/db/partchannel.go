package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
)

func PartChannel(target, channelName string, nb *bot.Bot) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := nb.MongoClient.Database("nourybot").Collection("channels")

	// Channel
	// {{"name": string}, {"connect": bool}}
	chnl := Channel{channelName, true}

	res, insertErr := collection.DeleteOne(ctx, chnl)

	// res.DeletedCount is 0 if trying to delete something that wasn't there.
	if insertErr != nil || res.DeletedCount != 1 {
		nb.Send(target, fmt.Sprintf("Error trying to part %s", channelName))
		log.Error(insertErr)
		return
	}
	nb.Send(target, fmt.Sprintf("Parted %s", channelName))

	// log.Info(res)
}
