package database

func (db *appdbimpl) CreateBan(user uint64, ban uint64) error {
	_, err := db.c.Exec(`INSERT INTO ban (user, ban) VALUES (?, ?)`, user, ban)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: ban.user, ban.ban" {
			// TO DO: con la put dovrei poter rimuovere questo check err.Error() perche non
			// dovrebbe essere possibile mettere like piu di una volta alla stessa immagine con lo stesso profilo
			return nil
		}
		return err
	}
	return nil
}
