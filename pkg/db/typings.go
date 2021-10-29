package db

type Channel struct {
	Name    string `bson:"name,omitempty"`
	Connect bool   `bson:"connect,omitempty"`
}

// type Command struct {
// 	Name    string `bson:"name,omitempty"`
// 	Text    string `bson:"text,omitempty"`
// 	Channel string `bson:"channel,omitempty"`
// }
