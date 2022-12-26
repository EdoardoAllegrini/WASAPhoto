package database

func (db *appdbimpl) GetPhotos(user uint64) ([]Photo, error) {
	rows, err := db.c.Query("SELECT id, users.username, caption, timestamp FROM media WHERE media.user=? and users.id=media.user ORDER BY timestamp DESC;", user)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []Photo
	// Here we read the resultset
	for rows.Next() {
		var tmp Photo
		err = rows.Scan(&tmp.ID, &tmp.User, &tmp.Caption, &tmp.Timestamp)
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
