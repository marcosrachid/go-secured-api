package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/marcosrachid/go-secured-api/internal/models/dto"
	"github.com/marcosrachid/go-secured-api/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI = "mongodb://mongo:mongo@localhost:27017"
	DATABASE  = "go-template"
)

func ConnectDB() *mongo.Collection {

	// Set client options
	clientOptions := options.Client().ApplyURI(utils.GetenvDefault("MONGO_URI", MONGO_URI))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database(utils.GetenvDefault("DATABASE", DATABASE)).Collection("books")

	return collection
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = dto.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
