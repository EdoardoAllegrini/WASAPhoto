package api

import (
	"encoding/json"
	"net/http"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the username from the request body.
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.IsValid() {
		// Here we validated the user structure content (username), and we
		// discovered that the username data is not valid.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Create the user in the database. Note that this function will return a new instance of the user with the
	// same information.
	dbnewuser, err := rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
	user.FromDatabase(dbnewuser)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user.Identifier)
}
