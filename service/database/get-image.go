package database

func (db *appdbimpl) GetImageFromIDPoster(photoid uint64, username string, uAuth string) ([]byte, error) {
	rows, err := db.c.Query(`SELECT image FROM media WHERE id=? and username=? and NOT EXISTS (select * from ban where ban.username=media.username and ban.ban=?);`, photoid, username, uAuth)

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
