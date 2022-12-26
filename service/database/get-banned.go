package database

func (db *appdbimpl) GetBanned(user uint64) ([]string, error) {
	rows, err := db.c.Query(`SELECT users.username FROM ban WHERE ban.user=? and ban.ban=users.id;`, user)

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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
