package bot

import (
	"time"

	twitch "github.com/gempir/go-twitch-irc/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bot struct {
	TwitchClient *twitch.Client
	MongoClient  *mongo.Client
	Uptime       time.Time
}
