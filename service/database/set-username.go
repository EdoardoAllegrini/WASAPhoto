package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetMyUserName(oldUser User, newusername string) (*User, error) {
	newidentifier := "ID_" + newusername
	_, err := db.c.Exec("UPDATE users SET username=?, identifier=? WHERE id=?", newusername, newidentifier, oldUser.ID)
	if err != nil {
		return nil, err
	}
	return &User{ID: oldUser.ID, Username: newusername, Identifier: newidentifier}, nil
}
