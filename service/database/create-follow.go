package database

func (db *appdbimpl) CreateFollow(username string, follow string) error {
	res, err := db.c.Exec(`INSERT INTO follow (username, follow) SELECT ?, ? WHERE NOT EXISTS (SELECT * FROM ban WHERE username=? and ban=?);`, username, follow, follow, username)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: follow.username, follow.follow" {
			// TO DO: con la put dovrei poter rimuovere questo check err.Error() perche non
			// dovrebbe essere possibile seguire piu di una volta la stessa persona
			return nil
		}
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrUserBanned
	}
	return nil
}
