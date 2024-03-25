package storage

import (
	"errors"

	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func dbconnection() (*sqlx.DB, error) {
	dbUri := os.Getenv("POSTGRES_URI")
	if dbUri == "" {
		// log.Fatal("POSTGRES_URI not found in environment variables")
		return nil, errors.New("POSTGRES_URI not found in environment variables")
	}
	db, err := sqlx.Connect("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
