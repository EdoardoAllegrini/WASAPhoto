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
	var u User
	u.Username = username
	// Check to avoid sql injection
	if !u.IsValid() {
		// Here we validated the user structure content (username), and we
		// discovered that the username data is not valid.
		w.WriteHeader(http.StatusNotFound)
		return
	}
	dbuser, err := rt.db.GetUserFromUsername(username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuser == nil {
		// The user does not exists.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get Authentication Token from Header
	auth_token := parseAuthToken(r)

	// Check if authentication is valid for username in path
	dbuserAuth, errAuth := rt.db.GetUserFromIdentifier(auth_token)
	if errAuth != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuserAuth == nil || dbuserAuth.Username != dbuser.Username {
		// Authentication not valid
		// User authenticated doesn't exist or doesn't matches user in path.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if username given in body is valid
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
