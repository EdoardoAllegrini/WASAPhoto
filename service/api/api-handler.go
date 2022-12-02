package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:username/", rt.wrap(rt.setMyUserName))
	rt.router.PUT("/users/:username/media/:photo-id/likes/:userLike/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/media/:photo-id/likes/:userLike/", rt.wrap(rt.unlikePhoto))

	// TO FIX: should be /:username/media
	rt.router.POST("/users/:username/media/", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/users/:username/media/:photo-id/", rt.wrap(rt.getImage))
	rt.router.DELETE("/users/:username/media/:photo-id/", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:username/media/:photo-id/likes/", rt.wrap(rt.getLikes))
	rt.router.PUT("/users/:username/following/:userFollow/", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/following/:userFollow/", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:username/following/", rt.wrap(rt.getFollowing))
	rt.router.GET("/users/:username/followers/", rt.wrap(rt.getFollowers))
	rt.router.PUT("/users/:username/ban/:userBan/", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/ban/:userBan/", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:username/ban/", rt.wrap(rt.getBanned))
	rt.router.POST("/users/:username/media/:photo-id/comments/", rt.wrap(rt.commentPhoto))
	rt.router.GET("/users/:username/media/:photo-id/comments/", rt.wrap(rt.getComments))
	rt.router.DELETE("/users/:username/media/:photo-id/comments/:comment-id/", rt.wrap(rt.uncommentPhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
