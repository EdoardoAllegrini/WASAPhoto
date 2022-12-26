package database

func (db *appdbimpl) GetImagePoster(photoid uint64) (string, error) {
	rows, err := db.c.Query(`SELECT users.username FROM media, users WHERE media.id=? and users.id=media.user;`, photoid)

	if err != nil {
		return "", err
	}
	defer func() { _ = rows.Close() }()

	var ret string
	// Here we read the resultset
	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			return "", err
		}
		ret = tmp
	}

	if rows.Err() != nil {
		return "", err
	}

	return ret, nil
}
