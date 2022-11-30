package database

func (db *appdbimpl) SetLike(photoid uint64, username uint64) error {
	_, err := db.c.Exec(`INSERT INTO likes (image, username) VALUES (?, ?);`, photoid, username)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: likes.image, likes.username" {
			// TO DO: con la put dovrei poter rimuovere questo check err.Error() perche non
			// dovrebbe essere possibile mettere like piu di una volta alla stessa immagine con lo stesso profilo
			return nil
		}
		return err
	}
	return nil
}
