package database

func (db *appdbimpl) SetLike(photoid uint64, poster string, username string) error {
	res, err := db.c.Exec(`INSERT INTO likes (image, username) SELECT ?, ? WHERE NOT EXISTS (SELECT * FROM ban WHERE ban.username=? and ban.ban=?);`, photoid, username, poster, username)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: likes.image, likes.username" {
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
