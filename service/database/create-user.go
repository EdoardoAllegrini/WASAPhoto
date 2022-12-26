package database

func (db *appdbimpl) CreateUser(u User) (User, error) {
	u.Identifier = "ID_" + u.Username

	res, err := db.c.Exec(`INSERT INTO users (username, identifier) VALUES (?, ?)`,
		u.Username, u.Identifier)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.identifier" || err.Error() == "UNIQUE constraint failed: users.username" {
			return u, nil
		}
		return u, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = uint64(lastInsertID)
	return u, nil
}
