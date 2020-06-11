package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/marcosrachid/go-secured-api/internal/db"
	"github.com/marcosrachid/go-secured-api/internal/models/repository"
	"github.com/marcosrachid/go-secured-api/pkg/response"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var books []repository.Book

	//Connection mongoDB with db class
	collection := db.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		response.GetError(err, w, http.StatusInternalServerError)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book repository.Book
		// & character returns the memory address of the following variable.
		err := cur.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		response.GetError(err, w, http.StatusInternalServerError)
	}

	response.GetResponse(books, w, http.StatusOK)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book repository.Book
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := db.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		response.GetError(err, w, http.StatusNotFound)
		return
	}

	response.GetResponse(book, w, http.StatusOK)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book repository.Book

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	// connect db
	collection := db.ConnectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		response.GetError(err, w, http.StatusInternalServerError)
		return
	}

	response.GetResponse(result, w, http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book repository.Book

	collection := db.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		response.GetError(err, w, http.StatusNotFound)
		return
	}

	book.ID = id

	response.GetResponse(book, w, http.StatusAccepted)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := db.ConnectDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		response.GetError(err, w, http.StatusNotFound)
		return
	}

	response.GetResponse(deleteResult, w, http.StatusAccepted)
}
