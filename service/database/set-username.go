package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetMyUserName(oldUser uint64, newusername string) (*User, error) {
	newidentifier := "ID_" + newusername
	_, err := db.c.Exec("UPDATE users SET username=?, identifier=? WHERE id=?", newusername, newidentifier, oldUser)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" || err.Error() == "UNIQUE constraint failed: users.identifier" {
			return nil, ErrUsernameNotAvailable
		}
		return nil, err
	}

	return &User{Username: newusername, Identifier: newidentifier}, nil
}
