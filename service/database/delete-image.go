package database

func (db *appdbimpl) DeleteImage(photoid uint64) ([]byte, error) {
	rows, err := db.c.Query(`DELETE FROM media WHERE id=?;`, photoid)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []byte
	// Here we read the resultset
	for rows.Next() {
		var tmp []byte
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		ret = tmp
	}

	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
