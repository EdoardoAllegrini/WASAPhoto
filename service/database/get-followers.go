package database

func (db *appdbimpl) GetFollowers(username uint64) ([]string, error) {
	rows, err := db.c.Query(`SELECT users.username FROM users, follow WHERE follow.follow=? and users.ID=follow.username;`, username)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []string
	// Here we read the resultset
	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
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
