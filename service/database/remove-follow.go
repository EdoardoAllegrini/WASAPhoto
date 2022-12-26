package database

func (db *appdbimpl) RemoveFollow(user uint64, follow uint64) error {
	res, err := db.c.Exec(`DELETE FROM follow WHERE user=? and follow=?;`, user, follow)
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
