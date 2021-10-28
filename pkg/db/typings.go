package db

type Channel struct {
	Name    string `bson:"name,omitempty"`
	Connect bool   `bson:"connect,omitempty"`
}
