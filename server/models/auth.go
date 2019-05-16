package models

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Credentials a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// User response from database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Claims a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

// JWTResponse struct returned from generate token
type JWTResponse struct {
	Token string
	Time  time.Time
}

// GetByUsername gets User by username
func GetByUsername(db *sql.DB, user *User, username string) error {
	err := db.QueryRow(
		`SELECT u.id, u.username, u.password FROM users u WHERE username = $1`,
		username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser returns User by username
func CreateUser(db *sql.DB, id *int, username string, password string) error {
	err := db.QueryRow(`INSERT INTO users(username, password)
		VALUES
		($1, $2)
		RETURNING id`, username, password).Scan(id)
	if err != nil {
		return err
	}
	return nil
}
