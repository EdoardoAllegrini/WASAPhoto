package api

import "WASAPhoto.uniroma1.it/wasaphoto/service/database"

func (rt *_router) checkAuth(auth_token string) (*database.User, error) {
	// Check if authentication is valid
	dbuser, err := rt.db.GetUserFromIdentifier(auth_token)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		return nil, err
	} else if dbuser == nil {
		// The user does not exists, authentication not valid.
		// Reject the action indicating an error on the client side.
		return nil, nil
	}
	return dbuser, nil
}
