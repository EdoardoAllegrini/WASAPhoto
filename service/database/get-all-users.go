package database

func (db *appdbimpl) GetAllUsers() ([]User, error) {
	rows, err := db.c.Query(`SELECT * FROM users;`)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []User
	// Here we read the resultset
	for rows.Next() {
		var tmp User
		err = rows.Scan(&tmp.ID, &tmp.Username, &tmp.Identifier)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
