package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func PartChannel(channelName string, mongoClient *mongo.Client) {
	// Interact with data

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/*
		Get my collection instance
	*/
	collection := mongoClient.Database("nourybot").Collection("channels")

	// Channel
	chnl := Channel{channelName, true}

	res, insertErr := collection.DeleteOne(ctx, chnl)
	if insertErr != nil {
		log.Error(insertErr)
	}
	log.Info(res)

}
