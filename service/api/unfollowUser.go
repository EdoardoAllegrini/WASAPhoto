package api

import (
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	userFollow := ps.ByName("userFollow")
	dbuserFo, errFo := rt.db.GetUserFromUsername(userFollow)
	if errFo != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuserFo == nil {
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

	// Check that user authenticated matches username given in path

	// The authentication and the user in path are valid.
	// Check if user authenticated matches to the one in path
	if dbuserAuth.Username != username {
		// User in path is different from the one authenticated
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("[+] Users are different")
		return
	}

	err = rt.db.RemoveFollow(dbuser.ID, dbuserFo.ID)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the user Auth")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.WriteHeader(http.StatusNoContent)
}
