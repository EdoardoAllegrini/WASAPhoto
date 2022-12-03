package database

func (db *appdbimpl) Comment(photoid uint64, username string, comment string) (uint64, error) {
	res, err := db.c.Exec(`INSERT INTO comments (image, username, comment) VALUES (?, ?, ?);`,
		photoid, username, comment)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertID), nil
}
