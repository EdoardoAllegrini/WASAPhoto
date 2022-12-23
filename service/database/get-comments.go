package database

func (db *appdbimpl) GetComments(photoid uint64, uAuth string) ([]Comment, error) {
	rows, err := db.c.Query(`SELECT comments.id, comments.image, comments.username, comment, comments.timestamp FROM comments WHERE comments.image=? and NOT EXISTS (SELECT * FROM ban WHERE ban.username=comments.username and ban.ban=?);`, photoid, uAuth)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []Comment
	// Here we read the resultset
	for rows.Next() {
		var tmp Comment
		err = rows.Scan(&tmp.ID, &tmp.Image, &tmp.User, &tmp.Text, &tmp.Timestamp)
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
