package database

func (db *appdbimpl) RemoveBan(user uint64, ban uint64) error {
	_, err := db.c.Exec(`DELETE FROM ban WHERE user=? and ban=?;`, user, ban)
	if err != nil {
		return err
	}
	return nil
}
