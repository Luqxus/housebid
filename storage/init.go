package storage

import (
	"database/sql"
	"log"
)

func dbconnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=yourdatabase sslmode=disable password=yourpassword host=localhost")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
