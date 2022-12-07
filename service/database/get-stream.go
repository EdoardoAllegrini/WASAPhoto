package database

import (
	"strconv"
)

func (db *appdbimpl) GetStream(user string) ([]Article, error) {
	qu := "SELECT media.id, media.username, media.caption, media.timestamp FROM media, follow WHERE follow.username=? and media.username=follow.follow and NOT EXISTS (select * from ban where ban.username=follow.follow and ban.ban=follow.username) ORDER BY timestamp DESC;"
	rows, err := db.c.Query(qu, user)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret []Article
	// Here we read the resultset
	for rows.Next() {
		var a Article
		var tmp Photo
		err = rows.Scan(&tmp.ID, &tmp.User, &tmp.Caption, &tmp.Timestamp)
		if err != nil {
			return nil, err
		}
		tmp.URL = "http://localhost:3000/users/" + tmp.URL + "/media/" + strconv.FormatUint(tmp.ID, 10) + "/"
		a.Ph = tmp
		likes, err := db.GetLikes(tmp.ID)
		if err != nil {
			return nil, err
		}
		a.Likes = len(likes)
		comments, err := db.GetComments(tmp.ID)
		if err != nil {
			return nil, err
		}
		a.Comments = len(comments)
		ret = append(ret, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
