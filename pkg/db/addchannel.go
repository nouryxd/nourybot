package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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

	/*
		Insert channel
	*/
	// chnl := []interface{}{
	// 	bson.D{{Key: "name", Value: "nouryqt"}, {Key: "connect", Value: true}},
	// 	bson.D{{Key: "name", Value: "nourybot"}, {Key: "connect", Value: true}},
	// }

	chnl := []interface{}{
		bson.D{{Key: "name", Value: channelName}, {Key: "connect", Value: true}},
	}

	res, insertErr := collection.InsertOne(ctx, chnl)
	if insertErr != nil {
		log.Error(insertErr)
	}
	log.Info(res)

}
