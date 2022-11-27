package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetMyUserName(oldUser User, newusername string) (*User, error) {
	newidentifier := "ID_" + newusername
	res, err := db.c.Exec("UPDATE users SET username=?, identifier=? WHERE id=?", newusername, newidentifier, oldUser.ID)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.identifier" {
			return nil, ErrUsernameNotAvailable
		}
		return nil, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return nil, ErrUserDoesNotExist
	}
	return &User{ID: oldUser.ID, Username: newusername, Identifier: newidentifier}, nil
}
