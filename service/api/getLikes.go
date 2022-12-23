package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Check if authentication is valid
	dbuserAuth, err := rt.db.GetUserFromIdentifier(auth_token)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuserAuth == nil {
		// The user does not exists, authentication not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if username in path has banned user authenticated
	// (done with separated query cause if condition on line up is true I can avoid to check that condition every query)
	c, errC := rt.db.CheckBanned(dbuser.Username, dbuserAuth.Username)
	if errC != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if c {
		// Username has banned user authenticated
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		// fmt.Println(dbuser.Username + " banned " + dbuserAuth.Username)
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
	// Query not unified in GetLikes because otherwise I can't distinguish between photo with no like and photo not in the db
	c, err = rt.db.CheckImagePoster(photoid, dbuser.Username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !c {
		// Photo with photo-id is not present in the db as a photo posted by username in path
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dbLikes, err := rt.db.GetLikes(photoid, dbuserAuth.Username)
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
	_ = json.NewEncoder(w).Encode(dbLikes)
}
