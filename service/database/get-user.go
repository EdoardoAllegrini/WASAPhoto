package database

func (db *appdbimpl) GetUser(u User) (*User, error) {
	rows, err := db.c.Query(`SELECT username, identifier FROM users WHERE username=? and identifier=?;`, u.Username, u.Identifier)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret *User
	// Here we read the resultset
	for rows.Next() {
		var tmp User
		err = rows.Scan(&tmp.Username, &tmp.Identifier)
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

func (db *appdbimpl) GetUserFromUsername(username string) (*User, error) {
	rows, err := db.c.Query(`SELECT username, identifier FROM users WHERE username=?;`, username)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret *User
	// Here we read the resultset
	for rows.Next() {
		var tmp User
		err = rows.Scan(&tmp.Username, &tmp.Identifier)
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

func (db *appdbimpl) GetUserFromIdentifier(identifier string) (*User, error) {
	rows, err := db.c.Query(`SELECT username, identifier FROM users WHERE identifier=?;`, identifier)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret *User
	// Here we read the resultset
	for rows.Next() {
		var tmp User
		err = rows.Scan(&tmp.Username, &tmp.Identifier)
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
