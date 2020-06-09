package main

import (
	"log"
	"net/http"

	"github.com/marcosrachid/go-secured-api/internal/api"
	"github.com/marcosrachid/go-secured-api/pkg/utils"

	"github.com/gorilla/mux"
)

const (
	PORT = "9090"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/books", api.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", api.GetBook).Methods("GET")
	r.HandleFunc("/api/books", api.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", api.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", api.DeleteBook).Methods("DELETE")

	// set our port address
	log.Printf("Serving at localhost:%s...\n", utils.GetenvDefault("PORT", PORT))
	log.Fatal(http.ListenAndServe(":"+utils.GetenvDefault("PORT", PORT), r))
}
