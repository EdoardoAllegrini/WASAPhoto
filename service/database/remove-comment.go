package database

func (db *appdbimpl) RemoveComment(commentid uint64, photoid uint64) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE id=? and image=?;`, commentid, photoid)
	if err != nil {
		return err
	}
	return nil
}
