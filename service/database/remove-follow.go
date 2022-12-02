package database

func (db *appdbimpl) RemoveFollow(username uint64, follow uint64) error {
	_, err := db.c.Exec(`DELETE FROM follow WHERE username=? and follow=?;`, username, follow)
	if err != nil {
		return err
	}
	return nil
}
