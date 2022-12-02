/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrUsernameNotAvailable = errors.New("username is not available")
var ErrImageDoesNotExist = errors.New("image does not exist")

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

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// CreateUser creates a new user in the database. It returns an updated User object (with the ID)
	CreateUser(User) (User, error)

	// GetUser returns the user in the database with id, username, identifier given
	GetUser(User) (*User, error)

	// GetUserFromUsername returns the user in the database with username given, if exists
	GetUserFromUsername(string) (*User, error)

	// GetUserFromIdentifier returns the user in the database with identifier given, if exists
	GetUserFromIdentifier(string) (*User, error)

	// SetMyUserName changes the username of User to string given in input
	SetMyUserName(User, string) (*User, error)

	// CreateMedia inserts the photo and its caption in the db
	CreateMedia(uint64, string, []byte) (uint64, error)

	// GetImageFromIDPoster returns the image in the database posted by username with id given, if exists
	GetImageFromIDPoster(uint64, uint64) ([]byte, error)

	// GetImagePoster returns the username of the user that posted photo with id given, if exists
	GetImagePoster(uint64) (string, error)

	// CheckImagePoster returns true if poster posted image, else false
	CheckImagePoster(uint64, uint64) (bool, error)

	// GetLikes returns the list of users that like the photo given in input
	GetLikes(uint64) ([]string, error)

	// DeleteImage deletes the image in the database with id given, if exists
	DeleteImage(uint64) error

	// SetLike is a function that let user relative to userLike set a like to the photo as photo-id posted by user given in path
	SetLike(uint64, uint64) error

	// SetLike is a function that let user relative to userLike remove the like to the photo as photo-id posted by user given in path
	RemoveLike(uint64, uint64) error

	// CreateFollow is a function that adds a new follower (userFollow) to username profile
	CreateFollow(uint64, uint64) error

	// RemoveFollow is a function that removes the follower (userFollow) to username profile
	RemoveFollow(uint64, uint64) error

	// GetFollowers returns the list of users followed by username
	GetFollowing(uint64) ([]string, error)

	// GetFollowers returns the list of users that follow username
	GetFollowers(uint64) ([]string, error)

	// CreateBan is a function that adds userFollow to users banned by username
	CreateBan(uint64, uint64) error

	// RemoveBan is a function that removes the banned userFollow to username banned users
	RemoveBan(uint64, uint64) error

	// GetBanned returns the list of users banned by username
	GetBanned(uint64) ([]string, error)

	// Comment adds a comment to the photo given, returns the id of the comment
	Comment(uint64, uint64, string) (uint64, error)

	// GetComments returns all the comment of the photo given
	GetComments(uint64) ([]Comment, error)

	// RemoveComment removes the comment given from the photo in path
	RemoveComment(uint64) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string

	// db.QueryRow(`DROP TABLE users;`).Scan(&tableName)
	// return nil, nil

	// db.QueryRow(`DROP TABLE likes;`).Scan(&tableName)
	// db.QueryRow(`DROP TABLE follow;`).Scan(&tableName)
	// db.QueryRow(`DROP TABLE ban;`).Scan(&tableName)
	// db.QueryRow(`DROP TABLE comments;`).Scan(&tableName)

	// return nil, nil

	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL, identifier TEXT NOT NULL UNIQUE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='media';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE media (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username INTEGER NOT NULL, caption TEXT NOT NULL, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, image BLOB NOT NULL, FOREIGN KEY (username) REFERENCES users(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (image INTEGER NOT NULL, username INTEGER NOT NULL, PRIMARY KEY (image, username), FOREIGN KEY (image) REFERENCES media(id) ON DELETE CASCADE, FOREIGN KEY (username) REFERENCES users(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='follow';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE follow (username INTEGER NOT NULL, follow INTEGER NOT NULL, PRIMARY KEY (username, follow), FOREIGN KEY (username) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY (follow) REFERENCES users(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='ban';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE ban (username INTEGER NOT NULL, ban INTEGER NOT NULL, PRIMARY KEY (username, ban), FOREIGN KEY (username) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY (ban) REFERENCES users(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, image INTEGER NOT NULL, username INTEGER NOT NULL, comment TEXT NOT NULL, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (image) REFERENCES media(id) ON DELETE CASCADE, FOREIGN KEY (username) REFERENCES users(id) ON DELETE CASCADE);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
