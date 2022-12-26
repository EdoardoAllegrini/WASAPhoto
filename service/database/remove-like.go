package database

func (db *appdbimpl) RemoveLike(photoid uint64, poster uint64, user uint64) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE image=? and user=? and NOT EXISTS (SELECT * FROM ban WHERE ban.user=? and ban.ban=?);`, photoid, user, poster, user)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// image doesn't exits or user that posted image banned user authenticated
		return ErrUserBanned
	}
	return nil
}
