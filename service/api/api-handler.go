package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	rt.router.PUT("/:username/", rt.wrap(rt.setMyUserName))

	// TO FIX: should be /:username/media
	rt.router.POST("/media/", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/media/:photo-id/", rt.wrap(rt.getImage))
	rt.router.DELETE("/media/:photo-id/", rt.wrap(rt.deletePhoto))
	rt.router.GET("/media/:photo-id/likes/", rt.wrap(rt.getLikes))
	// should be PUT but raises conflict error
	rt.router.POST("/media/:photo-id/likes/:userLike/", rt.wrap(rt.likePhoto))

	// TO DELETE
	rt.router.GET("/", rt.wrap(rt.getUsers))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
