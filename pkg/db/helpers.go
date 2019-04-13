package db

import (
	"database/sql"

	"github.com/reynld/go-auth/pkg/models"
)

// GetByUsername gets User by username
func GetByUsername(db *sql.DB, user *models.User, username string) error {
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
