package models

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"golang.org/x/crypto/bcrypt"
)

// RunMigrations runs migrations on database
func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

// RunDBSeeds migrates and seeds databse with JSON file from scraper
func RunDBSeeds(db *sql.DB) {
	hash, err := bcrypt.GenerateFromPassword([]byte("pass"), 10)
	if err != nil {
		log.Fatal("error hasing seed password")
	}

	var id int
	err = CreateUser(db, &id, "rey", string(hash))
	if err != nil {
		log.Fatal("error seeding user")
	}
}
