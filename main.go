package main

import (
	"database/sql"
	"fmt"
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

	// product := Product{"Book", 15.55, true}

	// primaryKey := insertProduct(db, product)

	// var name string
	// var available bool
	// var price float64

	// query := "SELECT name, available, price FROM product WHERE id= $1"
	// err := db.QueryRow(query, primaryKey).Scan(&name, &available, &price)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Fatal("No rows found with ID %d", primaryKey)
	// 	}
	// 	log.Fatal(error)
	// }

	// fmt.Printf("Name: %s\n", name)
	// fmt.Printf("Available: %t\n", available)
	// fmt.Printf("Price: %f\n", price)

	data := []Product{}
	query := "SELECT name, available, price FROM product"

	rows, err := db.Query(query)

	defer rows.Close() // prevent memory loss

	if err != nil {
		log.Fatal(err)
	}

	var name string
	var available bool
	var price float64

	for rows.Next() {
		err := rows.Scan(&name, &available, &price)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}

	fmt.Println(data)

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
