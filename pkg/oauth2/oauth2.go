package oauth2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/marcosrachid/go-secured-api/pkg/response"
	"github.com/marcosrachid/go-secured-api/pkg/utils"
)

type RealmAccess struct {
	RealmAccess Roles `json:"realm_access"`
}

type Roles struct {
	Roles []string `json:"roles"`
}

const (
	KEYCLOAK_AUTH_URL  = "http://localhost:8080/auth"
	KEYCLOAK_REALM     = "go"
	KEYCLOAK_CLIENT_ID = "go-sso"
)

func IsAuthorized(verifier *oidc.IDTokenVerifier, role string, endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		if rawAccessToken == "" {
			response.GetClientError(errors.New("Missing Token"), w, http.StatusUnauthorized)
			return
		}

		ctx := context.Background()
		token, err := verifier.Verify(ctx, rawAccessToken)

		if err != nil {
			log.Println(err.Error())
			response.GetClientError(errors.New("Unauthorized"), w, http.StatusUnauthorized)
			return
		}

		var realmAccess RealmAccess
		if err = token.Claims(&realmAccess); !utils.Contains(realmAccess.RealmAccess.Roles, role) {
			log.Printf("missing role: %s in %v\n", role, realmAccess)
			response.GetClientError(errors.New("Unauthorized"), w, http.StatusUnauthorized)
			return
		}

		endpoint(w, r)
	})
}

func GetResource() string {
	return fmt.Sprintf("%s/realms/%s", utils.GetenvDefault("KEYCLOAK_AUTH_URL", KEYCLOAK_AUTH_URL), utils.GetenvDefault("KEYCLOAK_REALM", KEYCLOAK_REALM))
}

func GetAuthUrl() string {
	return utils.GetenvDefault("KEYCLOAK_AUTH_URL", KEYCLOAK_AUTH_URL)
}

func GetRealm() string {
	return utils.GetenvDefault("KEYCLOAK_REALM", KEYCLOAK_REALM)
}

func GetClientId() string {
	return utils.GetenvDefault("KEYCLOAK_CLIENT_ID", KEYCLOAK_CLIENT_ID)
}
