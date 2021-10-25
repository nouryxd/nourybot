package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lyx0/nourybot/pkg/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	conf := config.LoadConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	// Interact with data
	type Channel struct {
		Name string `bson:"name,omitempty"`
	}

	/*
		Get my collection instance
	*/
	collection := client.Database("nourybot").Collection("channels")

	/*
		Insert channel
	*/
	// chnl := []interface{}{
	// 	bson.D{{"name", "nouryqt"}},
	// 	bson.D{{"name", "nourybot"}},
	// }

	// res, insertErr := collection.InsertMany(ctx, chnl)
	// if insertErr != nil {
	// 	log.Fatal(insertErr)
	// }
	// log.Info(res)

	/*
		Iterate a cursor
	*/

	// var result bson.M
	// err = collection.FindOne(ctx, bson.D{{}}).Decode(&result)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return
	// 	}

	// 	panic(err)
	// }
	// log.Info(result["name"])

	//------------------------------
	cur, currErr := collection.Find(ctx, bson.D{{}})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var results []bson.M
	if err = cur.All(ctx, &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result["name"])
	}
	//------------------------------
}

// https://docs.mongodb.com/drivers/go/current/fundamentals/crud/read-operations/retrieve/
