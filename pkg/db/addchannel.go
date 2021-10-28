package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
)

func AddChannel(target, channelName string, nb *bot.Bot) {
	// Interact with data

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/*
		Get my collection instance
	*/
	collection := nb.MongoClient.Database("nourybot").Collection("channels")

	// Channel
	// {{"name": string}, {"connect": bool}}
	chnl := Channel{channelName, true}

	res, insertErr := collection.InsertOne(ctx, chnl)
	if insertErr != nil {
		nb.Send(target, fmt.Sprintf("Error trying to join %s", channelName))
		log.Error(insertErr)
		return
	}
	nb.Send(target, fmt.Sprintf("Joined %s", channelName))
	log.Info(res)

}
