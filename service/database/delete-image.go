package database

func (db *appdbimpl) DeleteImage(photoid uint64) error {
	res, err := db.c.Exec(`DELETE FROM media WHERE id=?;`, photoid)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the image didn't exist
		return ErrImageDoesNotExist
	}

	return nil
}
