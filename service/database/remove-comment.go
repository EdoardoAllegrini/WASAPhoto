package database

func (db *appdbimpl) RemoveComment(commentid uint64) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE id=?;`, commentid)
	if err != nil {
		return err
	}
	return nil
}
