package api

import (
	"net/http"
	"strconv"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the username in path
	username := ps.ByName("username")

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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TO FIX: after fix in api-handler uncomment following 2 line and delete 3rd
	// photoid := ps.ByName("photo-id")
	photoid, err := strconv.ParseUint(ps.ByName("photo-id"), 10, 64)
	if err != nil {
		// Here we validated the photo-id given in path, and we
		// discovered that is not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("[-] Photo-id in path is not valid")
		return
	}
	c, err := rt.db.CheckImagePoster(photoid, dbuser.Username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !c {
		// Photo with photo-id is not present in the db as a photo posted by username in path
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
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
	} else if dbuserAuth == nil {
		// The user does not exists, authentication not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check that user authenticated matches userLike given in path

	// The authentication and the user in path are valid.
	// Check if user authenticated matches to the one in path
	if dbuserAuth.Username != userLike {
		// User in path is different from the one authenticated
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("[+] Users are different")
		return
	}

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

	errLike := rt.db.RemoveLike(photoid, dbuserAuth.Username)
	if errLike != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}