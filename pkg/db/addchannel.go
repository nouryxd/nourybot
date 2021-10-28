package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

func AddChannel(channelName string) {
	// Interact with data

	client := Connect()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/*
		Get my collection instance
	*/
	collection := client.Database("nourybot").Collection("channels")

	// Channel
	chnl := Channel{channelName, true}

	res, insertErr := collection.InsertOne(ctx, chnl)
	if insertErr != nil {
		log.Error(insertErr)
	}
	log.Info(res)

}
