package database

func (db *appdbimpl) CheckImagePoster(photoid uint64, username string) (uint64, error) {
	var ret User
	ret.ID = 0
	res, err := db.c.Query(`SELECT DISTINCT users.id FROM media, users WHERE media.id=? and users.username=? and media.user=users.id;`, photoid, username)

	if err != nil {
		return 0, err
	}
	defer func() { _ = res.Close() }()
	// Here we read the resultset
	for res.Next() {
		var u uint64
		err = res.Scan(&u)
		if err != nil {
			return 0, err
		}
		ret.ID = u
	}
	if err = res.Err(); err != nil {
		return 0, err
	}
	return ret.ID, nil
}
