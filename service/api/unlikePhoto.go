package api

import (
	"errors"
	"net/http"
	"strconv"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"WASAPhoto.uniroma1.it/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	photoid, err := strconv.ParseUint(ps.ByName("photo-id"), 10, 64)
	if err != nil {
		// Here we validated the photo-id given in path, and we
		// discovered that is not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("[-] Photo-id in path is not valid")
		return
	}

	userLike := ps.ByName("userLike")
	dbuserLiker, errLiker := rt.db.GetUserFromUsername(userLike)
	if errLiker != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuserLiker == nil {
		// The user does not exists.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get Authentication Token from Header
	auth_token := parseAuthToken(r)

	// Check if authentication is valid
	dbuserAuth, errAuth := rt.db.GetUserFromIdentifier(auth_token)
	if errAuth != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user Auth")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuserAuth == nil || dbuserAuth.ID != dbuserLiker.ID {
		// Authentication not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	errLike := rt.db.RemoveLike(photoid, dbuser.ID, dbuserAuth.ID)
	if errLike != nil {
		if errors.Is(errLike, database.ErrUserBanned) {
			// User to follow has banned username
			// Reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
