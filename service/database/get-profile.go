package database

import "strconv"

func (db *appdbimpl) GetProfile(userProf string) (*Profile, error) {
	rows, err := db.c.Query("SELECT media.id, media.username, media.caption, media.timestamp FROM media WHERE media.username=?  ORDER BY timestamp DESC;", userProf)

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
		tmp.URL = "/users/" + userProf + "/media/" + strconv.FormatUint(tmp.ID, 10) + "/"
		ret.Photos = append(ret.Photos, tmp)
		ret.N_Photos++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	followers, err := db.GetFollowers(userProf)
	if err != nil {
		return nil, err
	}
	ret.Followers = followers

	following, err := db.GetFollowing(userProf)
	if err != nil {
		return nil, err
	}
	ret.Following = following

	ret.User = userProf
	return &ret, nil
}
