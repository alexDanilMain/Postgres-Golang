package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

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

	createProductTable(db)
}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, error := db.Exec(query)

	if error != nil {
		log.Fatal(error)
	}

}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name,price,available)
	VALUES ($1, $2, $3) RETURNING id`

	var returnedId int
	error := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&returnedId)

	if error != nil {
		log.Fatal(error)
	}

	return returnedId
}
