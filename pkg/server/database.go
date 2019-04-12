package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// getDBKeys makes sure all enviroment variables are set and return them
func getDBKeys() (map[string]string, error) {
	keys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
	}

	values := map[string]string{}

	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return nil, fmt.Errorf("eviroment variable %s is required", key)
		}
		values[key] = v
	}

	return values, nil
}

// getDBUri makes sure all enviroment variables are set and return them
func getDBUri() (string, error) {
	d, err := getDBKeys()
	if err != nil {
		return "", err
	}

	var dburi string
	if d["ENVIROMENT"] == "PROD" {
		// Production
		dburi = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			d["DB_HOST"],
			d["DB_PORT"],
			d["DB_USER"],
			d["DB_NAME"],
			d["DB_PASSWORD"],
		)
	} else {
		// Local
		dburi = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			d["DB_HOST"],
			d["DB_PORT"],
			d["DB_USER"],
			d["DB_NAME"],
		)
	}

	return dburi, nil
}

// initializeDB connects to DB
func (s *Server) initializeDB() {
	dburi, err := getDBUri()
	if err != nil {
		log.Fatal(err)
	}

	s.DB, err = sql.Open("postgres", dburi)
	if err != nil {
		log.Fatal(err)
	}

	if err = s.DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

// RunMigrations runs migrations on database
func (s *Server) RunMigrations() {
	driver, err := postgres.WithInstance(s.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migration", "postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	// if err := m.Run(); err != nil && err != migrate.ErrNoChange {
	if err := m.Run(); err != nil {
		log.Fatal(err)
	}
}

// RunSeeds seeds DB with deafult user and password
func (s *Server) RunSeeds() {
	password := "pass"
	username := "rey"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf(`
	INSERT INTO users(username, password)
		VALUES
		('%s', '%s')
		RETURNING id
	`, username, string(hash))

	result, err := s.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
