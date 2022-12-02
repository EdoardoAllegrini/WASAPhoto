package api

import (
	"regexp"

	"WASAPhoto.uniroma1.it/wasaphoto/service/database"
)

const (
	UsernamePattern string = "^[a-zA-Z0-9._]+$"
)

type User struct {
	ID         uint64 `json:"id"`
	Username   string `json:"username"`
	Identifier string `json:"identifier"`
}

type Comment struct {
	ID        uint64 `json:"id"`
	Image     uint64 `json:"image"`
	User      uint64 `json:"username"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (u *User) FromDatabase(user database.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Identifier = user.Identifier
}

// ToDatabase returns the user in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:         u.ID,
		Username:   u.Username,
		Identifier: u.Identifier,
	}
}
func (u *User) IsValid() bool {
	r, _ := regexp.MatchString(UsernamePattern, u.Username)
	return r && len(u.Username) >= 3 && len(u.Username) <= 16
}
