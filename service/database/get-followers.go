package database

func (db *appdbimpl) GetFollowers(user uint64, uAuth uint64) ([]string, error) {
	rows, err := db.c.Query(`SELECT users.username FROM follow, users WHERE follow.follow=? and NOT EXISTS (select * from ban where ban.user=follow.user and ban.ban=?) and follow.user=users.id;`, user, uAuth)

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
