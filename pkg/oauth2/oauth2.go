package oauth2

import (
	"fmt"
	"net/http"

	"github.com/marcosrachid/go-secured-api/pkg/utils"
)

const (
	KEYCLOAK_AUTH_URL         = "http://localhost:8080/auth"
	KEYCLOAK_REALM            = "go"
	KEYCLOAK_ADAPTER_RESOURCE = "go-sso"
)

func IsAuthorized(provider string, endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		endpoint(w, r)
	})
}

func GetResource() string {
	return fmt.Sprintf("%s/realms/%s/protocol/openid-connect/", utils.GetenvDefault("KEYCLOAK_AUTH_URL", KEYCLOAK_AUTH_URL), utils.GetenvDefault("KEYCLOAK_REALM", KEYCLOAK_REALM))
}
