package db

import (
	"context"
	"fmt"
	"log"

	"github.com/marcosrachid/go-secured-api/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI = "mongodb://mongo:mongo@localhost:27017"
	DATABASE  = "go-template"
)

var collection *mongo.Collection

func ConnectDB() *mongo.Collection {

	if collection == nil {
		// Set client options
		clientOptions := options.Client().ApplyURI(utils.GetenvDefault("MONGO_URI", MONGO_URI))

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")

		collection = client.Database(utils.GetenvDefault("DATABASE", DATABASE)).Collection("books")
	}

	return collection
}
