package api

import (
	"encoding/json"
	"io"
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func isIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

const MaxMemory int64 = 500000

func getSupportedFiles() []string {
	return []string{"image/png", "image/jpeg", "image/jpg"}
}
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
		w.WriteHeader(http.StatusNotFound)
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
	} else if dbuser == nil || dbuser.ID != dbuserPath.ID {
		// Authentication not valid.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(MaxMemory)
	if err != nil {
		// Request not parsable
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photoCaption := r.FormValue("photoCaption")
	if len(photoCaption) > 250 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photoFile, hea, err := r.FormFile("photoFile")

	if err != nil {
		// The image file is not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer photoFile.Close()
	fileType := hea.Header.Get("Content-Type")
	if !isIn(fileType, getSupportedFiles()) {
		// The image file is not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// fmt.Println(filepath.Ext(fileHeader.Filename))
	// read image got into a byte array
	fileBytes, err := io.ReadAll(photoFile)
	if err != nil {
		// The image file is not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	imageID, err := rt.db.CreateMedia(dbuser.ID, photoCaption, fileBytes)
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
