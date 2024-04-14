package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"

	db, error := sql.Open("postgres", connectionString)

	defer db.Close() // ensure db is closed'

	if error != nil {
		log.Fatal(error)
	}

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}
}
