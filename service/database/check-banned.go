package database

func (db *appdbimpl) CheckBanned(username string, userBanned string) (bool, error) {
	check, err := db.c.Query(`SELECT * FROM ban WHERE username=? and ban=?;`, username, userBanned)

	if err != nil {
		return false, err
	}
	defer func() { _ = check.Close() }()

	return check.Next(), nil
}