package database

func (db *appdbimpl) CreateMedia(username uint64, caption string, image []byte) (uint64, error) {
	res, err := db.c.Exec(`INSERT INTO media (username, caption, image) VALUES (?, ?, ?)`,
		username, caption, image)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertID), nil
}
