package database

func (db *appdbimpl) GetFollowers(username string) ([]string, error) {
	rows, err := db.c.Query(`SELECT follow.username FROM follow WHERE follow.follow=?;`, username)

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
