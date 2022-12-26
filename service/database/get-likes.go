package database

func (db *appdbimpl) GetLikes(photoid uint64, uAuth uint64) ([]string, error) {
	rows, err := db.c.Query(`SELECT users.username FROM likes, users WHERE likes.image=? and NOT EXISTS (SELECT * FROM ban WHERE ban.user=likes.user and ban.ban=?) and likes.user=users.id`, photoid, uAuth)

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
