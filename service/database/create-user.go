package database

func (db *appdbimpl) CreateUser(u User) (User, error) {
	u.Identifier = "ID_" + u.Username
	// res, err := db.c.Exec(`INSERT INTO users (username, identifier) VALUES (?, ?)`,

	res, err := db.c.Exec(`INSERT INTO users (username, identifier) VALUES (?, ?)`,
		u.Username, u.Identifier)
	if err != nil {
		return u, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.ID = uint64(lastInsertID)
	// fmt.Printf("[+] Created id: %d, username: %s, identifier: %s\n", u.ID, u.Username, u.Identifier)
	return u, nil
}
