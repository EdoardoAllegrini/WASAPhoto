package database

func (db *appdbimpl) GetBanned(username uint64) ([]string, error) {
	rows, err := db.c.Query(`SELECT users.username FROM users, ban WHERE ban.username=? and users.ID=ban.ban;`, username)

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
