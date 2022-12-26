package database

func (db *appdbimpl) CreateFollow(user uint64, follow uint64) error {
	res, err := db.c.Exec(`INSERT INTO follow (user, follow) SELECT ?, ? WHERE NOT EXISTS (SELECT * FROM ban WHERE ban.user=? and ban.ban=?);`, user, follow, follow, user)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: follow.user, follow.follow" {
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
