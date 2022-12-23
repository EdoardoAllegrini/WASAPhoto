package database

func (db *appdbimpl) GetFollowers(username string, uAuth string) ([]string, error) {
	rows, err := db.c.Query(`SELECT follow.username FROM follow WHERE follow.follow=? and NOT EXISTS (select * from ban where ban.username=follow.username and ban.ban=?);`, username, uAuth)

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
