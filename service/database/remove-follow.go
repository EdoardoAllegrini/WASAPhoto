package database

func (db *appdbimpl) RemoveFollow(username string, follow string) error {
	res, err := db.c.Exec(`DELETE FROM follow WHERE username=? and follow=?;`, username, follow)
	if err != nil {
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
