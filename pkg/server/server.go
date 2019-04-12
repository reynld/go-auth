package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

// GetDBKeys makes sure all enviroment variables are set and return them
func GetDBKeys() (map[string]string, error) {
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

// GetDBUri makes sure all enviroment variables are set and return them
func GetDBUri() (string, error) {
	d, err := GetDBKeys()
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

// RunMigrations runs migrations on database
func (s *Server) RunMigrations() {
	// Run migrations
	driver, err := postgres.WithInstance(s.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Run(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

// initializeDB connects to DB
func (s *Server) initializeDB() {
	dburi, err := GetDBUri()
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

// Initialize Server struct
func (s *Server) Initialize() {
	s.initializeDB()
	s.initializeRoutes()
}

// Run runs the router
func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
