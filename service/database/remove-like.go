package database

func (db *appdbimpl) RemoveLike(photoid uint64, username string) error {
	_, err := db.c.Exec(`DELETE FROM likes WHERE image=? and username=?;`, photoid, username)
	if err != nil {
		return err
	}
	return nil
}
