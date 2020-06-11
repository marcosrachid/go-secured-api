package main

import (
	"log"
	"net/http"

	"github.com/marcosrachid/go-secured-api/internal/api"
	"github.com/marcosrachid/go-secured-api/pkg/oauth2"
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
	r.Handle("/api/books", oauth2.IsAuthorized("list-books", api.GetBooks)).Methods("GET")
	r.Handle("/api/books/{id}", oauth2.IsAuthorized("get-book", api.GetBook)).Methods("GET")
	r.Handle("/api/books", oauth2.IsAuthorized("create-book", api.CreateBook)).Methods("POST")
	r.Handle("/api/books/{id}", oauth2.IsAuthorized("update-book", api.UpdateBook)).Methods("PUT")
	r.Handle("/api/books/{id}", oauth2.IsAuthorized("delete-book", api.DeleteBook)).Methods("DELETE")

	// set our port address
	log.Printf("Serving at localhost:%s...\n", utils.GetenvDefault("PORT", PORT))
	log.Fatal(http.ListenAndServe(":"+utils.GetenvDefault("PORT", PORT), r))
}
