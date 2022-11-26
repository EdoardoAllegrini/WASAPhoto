package api

import (
	"net/http"
	"strings"
)

func parseAuthToken(r *http.Request) string {
	// Parse the authentication token got
	auth_token := r.Header.Get("Authorization")
	auth_token = strings.Split(auth_token, "Bearer ")[1]
	return auth_token
}
