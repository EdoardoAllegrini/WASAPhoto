package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get Authentication Token from Header
	auth_token := parseAuthToken(r)
	fmt.Printf("[+] Got %s as token\n", auth_token)
	// Get the username in path
	username := ps.ByName("username")
	fmt.Printf("[+] Got %s in path\n", username)

	// Check if authentication is valid
	dbuser, err := rt.db.GetUserFromIdentifier(auth_token)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuser == nil {
		// The user does not exists, authentication not valid
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get User relative to username if exists
	var user User
	user.Username = username
	if !user.IsValid() {
		// Here we validated the user structure content (username), and we
		// discovered that the username data is not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("[-] Username in path is not valid")
		return
	}

	// The authentication and the user in path are valid.
	// Check if user authenticated corresponds to the one in path
	if dbuser.Username != user.Username {
		// User in path is different from the one authenticated
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("[+] Users are different")
		return
	}
	fmt.Println("[+] Users are the same")
	// If username given in body is valid and available the sets it
	var userBody User
	err = json.NewDecoder(r.Body).Decode(&userBody)
	_ = r.Body.Close()
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !userBody.IsValid() {
		// Here we validated the user structure content (username), and we
		// discovered that the username data is not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("[-] Username in body is not valid")
		return
	} else if userBody.Username == username {
		// If the username given in body is already the username of the user authenticated then sends something
		// TO FIX: http request deve dire che l'utente ha già questo username
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("[-] The user authenticated %s has already this username\n", auth_token)
		return
	}
	// Checks if the username given in body is available
	dbBodyuser, _ := rt.db.GetUserFromUsername(userBody.Username)
	if dbBodyuser != nil {
		// username is not available
		w.WriteHeader(http.StatusConflict)
		return
	}
	fmt.Println("[+] Username in body available")
	// username is available, changes it
	dbNewUser, err := rt.db.SetMyUserName(*dbuser, userBody.Username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(dbNewUser.Identifier)

}
