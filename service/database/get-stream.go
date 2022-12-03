package database

import (
	"fmt"
	"strings"
)

func (db *appdbimpl) GetStream(followed []string) ([]Article, error) {
	rows, err := db.c.Query("SELECT id FROM users WHERE username IN (?) ORDER BY timestamp DESC;", strings.Join(followed, ", "))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []Photo
	// Here we read the resultset
	for rows.Next() {
		var tmp Photo
		err = rows.Scan(&tmp.ID, &tmp.User, &tmp.Caption, &tmp.Timestamp)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(ret)
	return nil, nil
}
