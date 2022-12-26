package database

func (db *appdbimpl) GetImageFromIDPoster(photoid uint64, user uint64, uAuth uint64) ([]byte, error) {
	rows, err := db.c.Query(`SELECT image FROM media WHERE id=? and user=? and NOT EXISTS (select * from ban where ban.user=media.user and ban.ban=?);`, photoid, user, uAuth)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []byte
	// Here we read the resultset
	for rows.Next() {
		var tmp []byte
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		ret = tmp
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
