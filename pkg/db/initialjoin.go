package db

import (
	"context"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func InitialJoin(nb *bot.Bot) {
	collection := nb.MongoClient.Database("nourybot").Collection("channels")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// defer nb.MongoClient.Disconnect(ctx)

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var channels []Channel
	if err := cur.All(ctx, &channels); err != nil {
		panic(err)
	}

	for _, ch := range channels {
		nb.TwitchClient.Join(ch.Name)
		nb.TwitchClient.Say(ch.Name, "xd")
		log.Infof("Joined: %s\n", ch.Name)
	}

}
