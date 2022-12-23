package database

func (db *appdbimpl) Comment(phPoster string, photoid uint64, username string, comment string) (uint64, error) {
	res, err := db.c.Exec(`INSERT INTO comments (image, username, comment) SELECT ?, ?, ? WHERE NOT EXISTS (SELECT * FROM ban WHERE username=? and ban=?);`,
		photoid, username, comment, phPoster, username)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else if affected == 0 {
		return 0, nil
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertID), nil
}
