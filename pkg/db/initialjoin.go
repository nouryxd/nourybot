package db

import (
	"context"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	"go.mongodb.org/mongo-driver/bson"
)

func InitialJoin(nb *bot.Bot) {
	client := Connect()

	collection := client.Database("nourybot").Collection("channels")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer client.Disconnect(ctx)

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var channels []channel
	if err := cur.All(ctx, &channels); err != nil {
		panic(err)
	}

	for _, ch := range channels {
		nb.TwitchClient.Join(ch.Name)
		nb.TwitchClient.Say(ch.Name, "xd")
		// fmt.Printf("%v\n", ch.Name)
	}

}
