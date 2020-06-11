package main

import (
	"context"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/marcosrachid/go-secured-api/internal/api"
	"github.com/marcosrachid/go-secured-api/pkg/oauth2"
	"github.com/marcosrachid/go-secured-api/pkg/utils"

	"github.com/gorilla/mux"
)

const (
	PORT = "9090"
)

func main() {
	// Oauth2 provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, oauth2.GetResource())
	if err != nil {
		panic(err)
	}

	oidcConfig := &oidc.Config{
		ClientID: oauth2.GetClientId(),
	}
	verifier := provider.Verifier(oidcConfig)

	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.Handle("/api/books", oauth2.IsAuthorized(verifier, "list-books", api.GetBooks)).Methods("GET")
	r.Handle("/api/books/{id}", oauth2.IsAuthorized(verifier, "get-book", api.GetBook)).Methods("GET")
	r.Handle("/api/books", oauth2.IsAuthorized(verifier, "create-book", api.CreateBook)).Methods("POST")
	r.Handle("/api/books/{id}", oauth2.IsAuthorized(verifier, "update-book", api.UpdateBook)).Methods("PUT")
	r.Handle("/api/books/{id}", oauth2.IsAuthorized(verifier, "delete-book", api.DeleteBook)).Methods("DELETE")

	// set our port address
	log.Printf("Serving at localhost:%s...\n", utils.GetenvDefault("PORT", PORT))
	log.Fatal(http.ListenAndServe(":"+utils.GetenvDefault("PORT", PORT), r))
}
