package database

func (db *appdbimpl) Comment(phPoster uint64, photoid uint64, user uint64, comment string) (uint64, error) {
	res, err := db.c.Exec(`INSERT INTO comments (image, user, comment) SELECT ?, ?, ? WHERE NOT EXISTS (SELECT * FROM ban WHERE ban.user=? and ban.ban=?);`,
		photoid, user, comment, phPoster, user)
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
