package database

func (db *appdbimpl) CheckBanned(user uint64, userBanned uint64) (bool, error) {
	check, err := db.c.Query(`SELECT * FROM ban WHERE user=? and ban=?;`, user, userBanned)

	if err != nil {
		return false, err
	}
	defer func() { _ = check.Close() }()

	if err = check.Err(); err != nil {
		return false, err
	}

	return check.Next(), nil
}
