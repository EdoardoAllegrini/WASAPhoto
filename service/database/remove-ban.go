package database

func (db *appdbimpl) RemoveBan(username string, ban string) error {
	_, err := db.c.Exec(`DELETE FROM ban WHERE username=? and ban=?;`, username, ban)
	if err != nil {
		return err
	}
	return nil
}
