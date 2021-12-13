package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lyx0/nourybot/cmd/bot"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// InitialJoin is called everytime the Bot starts and joins the
// initial list of twitch channels it should be in.
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

	channelCount := 0

	for _, ch := range channels {
		nb.TwitchClient.Join(ch.Name)
		// nb.TwitchClient.Say(ch.Name, "xd")
		channelCount++
		log.Infof("Joined: %s\n", ch.Name)
	}
	// It worked
	nb.Send("nourybot", fmt.Sprintf("Badabing Badaboom Pepepains Joined %v channel", channelCount))
	nb.Send("nouryxd", fmt.Sprintf("Badabing Badaboom Pepepains Joined %v channel", channelCount))
}
