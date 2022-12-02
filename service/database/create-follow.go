package database

func (db *appdbimpl) CreateFollow(username uint64, follow uint64) error {
	_, err := db.c.Exec(`INSERT INTO follow (username, follow) VALUES (?, ?)`, username, follow)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: follow.username, follow.follow" {
			// TO DO: con la put dovrei poter rimuovere questo check err.Error() perche non
			// dovrebbe essere possibile mettere like piu di una volta alla stessa immagine con lo stesso profilo
			return nil
		}
		return err
	}
	return nil
}
