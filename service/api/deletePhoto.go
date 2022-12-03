package api

import (
	"errors"
	"net/http"
	"strconv"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"WASAPhoto.uniroma1.it/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// TO FIX: after fix in api-handler uncomment following 2 line and delete 3rd
	// Get the username in path
	username := ps.ByName("username")

	photoid, err := strconv.ParseUint(ps.ByName("photo-id"), 10, 64)
	if err != nil {
		// Here we validated the photo-id given in path, and we
		// discovered that is not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("[-] Photo-id in path is not valid")
		return
	}
	c, err := rt.db.CheckImagePoster(photoid, username)
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

	if username != dbuserAuth.Username {
		// The user authenticated can't delete the image with id given in path
		// because isn't the poster
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	errDel := rt.db.DeleteImage(photoid)
	if errors.Is(errDel, database.ErrImageDoesNotExist) {
		// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errDel != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't delete the image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusNoContent)
}
