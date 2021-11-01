package db

import (
	"context"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	"go.mongodb.org/mongo-driver/bson"
)

// Listchannel reads the list of channel from the database and
// whispers them to my own account.
func ListChannel(nb *bot.Bot) {
	collection := nb.MongoClient.Database("nourybot").Collection("channels")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var channels []Channel
	if err := cur.All(ctx, &channels); err != nil {
		panic(err)
	}

	channelList := ""
	for _, ch := range channels {
		channelList += ch.Name + " "
	}

	nb.TwitchClient.Whisper("nouryqt", channelList)
}
