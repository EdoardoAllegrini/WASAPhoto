package database

import "fmt"

func (db *appdbimpl) GetComments(photoid uint64) ([]Comment, error) {
	rows, err := db.c.Query(`SELECT id, image, username, comment, timestamp FROM comments WHERE image=?;`, photoid)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []Comment
	// Here we read the resultset
	for rows.Next() {
		var tmp Comment
		err = rows.Scan(&tmp.ID, &tmp.Image, &tmp.User, &tmp.Text, &tmp.Timestamp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ret = append(ret, tmp)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
