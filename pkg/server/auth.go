package server

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Signin the Signin handler
func (s *Server) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User

	err = s.DB.QueryRow(
		`SELECT u.id, u.username, u.password FROM users u WHERE username = $1`,
		creds.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtToken, err := GenerateToken(user.Username, user.ID)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtToken.Token,
		Expires: jwtToken.Time,
	})
}
