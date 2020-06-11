package oauth2

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/marcosrachid/go-secured-api/pkg/response"
	"github.com/marcosrachid/go-secured-api/pkg/utils"
)

type KeycloakCert struct {
	Keys []KeycloakCert `json:"keys"`
}

type KeycloakKey struct {
	Kid     string   `json:"kid"`
	Kty     string   `json:"kty"`
	Alg     string   `json:"alg"`
	Use     string   `json:"use"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	X5c     []string `json:"x5c"`
	X5t     string   `json:"x5t"`
	X5tS256 string   `json:"x5t#S256"`
}

const (
	KEYCLOAK_AUTH_URL         = "http://localhost:8080/auth"
	KEYCLOAK_REALM            = "go"
	KEYCLOAK_ADAPTER_RESOURCE = "go-sso"
)

func IsAuthorized(role string, endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return nil, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			response.GetError(errors.New("Unauthorized"), w, http.StatusUnauthorized)
		}
	})
}

func GetResource() string {
	return fmt.Sprintf("%s/realms/%s/protocol/openid-connect/", utils.GetenvDefault("KEYCLOAK_AUTH_URL", KEYCLOAK_AUTH_URL), utils.GetenvDefault("KEYCLOAK_REALM", KEYCLOAK_REALM))
}
