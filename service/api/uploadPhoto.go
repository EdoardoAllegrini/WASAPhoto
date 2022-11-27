package api

import (
	"fmt"
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

const MaxMemory int64 = 14000

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	r.ParseMultipartForm(MaxMemory)
	photoCaption := r.PostFormValue("photoCaption")
	photoFile := r.PostFormValue("photoFile")

	fmt.Printf("Caption : %s\n", photoCaption)
	fmt.Printf("Photo : %s\n", photoFile)
	return
}
