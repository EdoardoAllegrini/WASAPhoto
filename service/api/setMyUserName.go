package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"WASAPhoto.uniroma1.it/wasaphoto/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the username in path
	username := ps.ByName("username")

	// Get User relative to username given in path if exists
	var user User
	user.Username = username
	if !user.IsValid() {
		// Here we validated the user structure content (username), and we
		// discovered that the username data is not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("[-] Username in path is not valid")
		return
	}

	// Get Authentication Token from Header
	auth_token := parseAuthToken(r)

	// Check if authentication is valid
	dbuser, err := rt.db.GetUserFromIdentifier(auth_token)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuser == nil {
		// The user does not exists, authentication not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// The authentication and the user in path are valid.
	// Check if user authenticated matches to the one in path
	if dbuser.Username != user.Username {
		// User in path is different from the one authenticated
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("[+] Users are different")
		return
	}
	// Check if username given in body is valid and available
	var userBody User
	err = json.NewDecoder(r.Body).Decode(&userBody)

	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !userBody.IsValid() {
		// Here we validated the user structure content (username), and we
		// discovered that the username data is not valid.
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("[-] Username in body is not valid")
		return
	}

	dbNewUser, err := rt.db.SetMyUserName(*dbuser, userBody.Username)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		// The username in path does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if errors.Is(err, database.ErrUsernameNotAvailable) {
		// The username in body is not available, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusConflict)
		// fmt.Printf("[-] Username %s ain't available\n", userBody.Username)
		return
	}
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
