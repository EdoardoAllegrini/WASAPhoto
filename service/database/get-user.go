package database

func (db *appdbimpl) GetUser(u User) (*User, error) {
	rows, err := db.c.Query(`SELECT id, username, identifier FROM users WHERE username=?;`, u.Username)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret *User
	// Here we read the resultset
	for rows.Next() {
		var tmp User
		err = rows.Scan(&tmp.ID, &tmp.Username, &tmp.Identifier)
		if err != nil {
			return nil, err
		}
		ret = &tmp
	}

	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
