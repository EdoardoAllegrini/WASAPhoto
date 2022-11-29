package database

func (db *appdbimpl) CheckImagePoster(photoid uint64, username uint64) (bool, error) {
	check, err := db.c.Query(`SELECT id FROM media WHERE id=? and username=?;`, photoid, username)

	if err != nil {
		return false, err
	}
	defer func() { _ = check.Close() }()

	return check.Next(), nil
}
