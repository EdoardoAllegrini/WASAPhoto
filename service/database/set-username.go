package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetMyUserName(oldUser User, newusername string) (*User, error) {
	newidentifier := "ID_" + newusername
	_, err := db.c.Exec("UPDATE users SET username=?, identifier=? WHERE username=?", newusername, newidentifier, oldUser.Username)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			return nil, ErrUsernameNotAvailable
		}
		return nil, err
	}

	return &User{Username: newusername, Identifier: newidentifier}, nil
}
