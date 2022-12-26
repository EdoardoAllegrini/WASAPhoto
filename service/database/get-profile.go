package database

import "strconv"

func (db *appdbimpl) GetProfile(userProf uint64, uAuth uint64) (*Profile, error) {
	rows, err := db.c.Query("SELECT media.id, users.username, media.caption, media.timestamp FROM media, users WHERE media.user=? and media.user=users.id ORDER BY timestamp DESC;", userProf)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ret Profile
	// Here we read the resultset
	for rows.Next() {
		var tmp Photo
		err = rows.Scan(&tmp.ID, &tmp.User, &tmp.Caption, &tmp.Timestamp)
		if err != nil {
			return nil, err
		}
		tmp.URL = "/users/" + tmp.User + "/media/" + strconv.FormatUint(tmp.ID, 10) + "/"
		ret.Photos = append(ret.Photos, tmp)
		ret.N_Photos++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	followers, err := db.GetFollowers(userProf, uAuth)
	if err != nil {
		return nil, err
	}
	ret.Followers = followers

	following, err := db.GetFollowing(userProf, uAuth)
	if err != nil {
		return nil, err
	}
	ret.Following = following

	return &ret, nil
}
