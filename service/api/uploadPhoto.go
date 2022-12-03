package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

const MaxMemory int64 = 1000000

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the username in path
	username := ps.ByName("username")

	// Get User relative to username given in path if exists
	dbuserPath, errPa := rt.db.GetUserFromUsername(username)
	if errPa != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(errPa).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbuserPath == nil {
		// The user does not exists, authentication not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
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
	if dbuser.Username != username {
		// User in path is different from the one authenticated
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("[+] Users are different")
		return
	}
	r.ParseMultipartForm(MaxMemory)
	photoCaption := r.FormValue("photoCaption")
	photoFile, _, err := r.FormFile("photoFile")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer photoFile.Close()
	// fmt.Println(filepath.Ext(fileHeader.Filename))
	// read image got into a byte array
	fileBytes, err := io.ReadAll(photoFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	imageID, err := rt.db.CreateMedia(dbuser.Username, photoCaption, fileBytes)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't store image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(imageID)
}
