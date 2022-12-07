package database

func (db *appdbimpl) CreateUser(u User) (User, error) {
	u.Identifier = "ID_" + u.Username

	_, err := db.c.Exec(`INSERT INTO users (username, identifier) VALUES (?, ?)`,
		u.Username, u.Identifier)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.identifier" {
			return u, nil
		}
		return u, err
	}
	return u, nil
}
