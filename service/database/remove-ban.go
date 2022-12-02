package database

func (db *appdbimpl) RemoveBan(username uint64, ban uint64) error {
	_, err := db.c.Exec(`DELETE FROM ban WHERE username=? and ban=?;`, username, ban)
	if err != nil {
		return err
	}
	return nil
}
