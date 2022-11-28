package api

import (
	"net/http"
	"strconv"

	"WASAPhoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	dbImage, err := rt.db.GetImageFromID(photoid)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get the image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbImage == nil {
		// The image does not exists.
		// Reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(dbImage)
}
