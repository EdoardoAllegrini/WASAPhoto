package database

func (db *appdbimpl) RemoveLike(photoid uint64, poster string, username string) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE image=? and username=? and NOT EXISTS (SELECT * FROM ban WHERE ban.username=? and ban.ban=?);`, photoid, username, poster, username)
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
