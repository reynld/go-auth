package server

import (
	"fmt"
	"os"
)

// CheckEnviroment function makes sure all enviroment variables are set befire running
func (s *Server) CheckEnviroment() error {
	variables := []string{
		"PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
		"DB_PORT",
		"ENVIROMENT",
		"JWT_KEY",
		"CACHE_ADDRS",
		"CACHE_DB",
	}

	for _, v := range variables {
		if value := os.Getenv(v); value == "" {
			return fmt.Errorf("enviroment variable %s is required", v)
		}
	}

	return nil
}

// Panic checks if error exists
func (s *Server) Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// GetDBKeys makes sure all enviroment variables are set and return them
func (s *Server) GetDBKeys() (map[string]string, error) {
	keys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
		"ENVIROMENT",
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
func (s *Server) GetDBUri() (string, error) {
	d, err := s.GetDBKeys()
	if err != nil {
		return "", err
	}

	var dburi string
	if d["ENVIROMENT"] == "PRO" {
		// Production
		dburi = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			d["DB_HOST"],
			d["DB_PORT"],
			d["DB_USER"],
			d["DB_NAME"],
			d["DB_PASSWORD"],
		)
	} else if d["ENVIROMENT"] == "DEV" {
		// Local
		dburi = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			d["DB_HOST"],
			d["DB_PORT"],
			d["DB_USER"],
			d["DB_NAME"],
		)
	} else {
		return "", fmt.Errorf("ENVIROMENT variable must be PRO or DEV")
	}

	return dburi, nil
}
