package database

func (db *appdbimpl) SetLike(photoid uint64, poster uint64, user uint64) error {
	res, err := db.c.Exec(`INSERT INTO likes (image, user) SELECT ?, ? WHERE NOT EXISTS (SELECT * FROM ban WHERE ban.user=? and ban.ban=?);`, photoid, user, poster, user)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: likes.image, likes.user" {
			// TO DO: con la put dovrei poter rimuovere questo check err.Error() perche non
			// dovrebbe essere possibile mettere like piu di una volta alla stessa immagine con lo stesso profilo
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
