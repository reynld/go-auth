package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// User response from database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Credentials a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
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
