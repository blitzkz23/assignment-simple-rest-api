package database

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	username = "postgres"
	password = "naufalaldy23"
	host     = "localhost"
	dbPort   = "5432"
	dbName   = "assignment2"
	dbDriver = "postgres"
	db       *sql.DB
	err      error
)

func createRequiredTable() {
	// ! SQL Stament to create order and item table
	orderTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_name VARCHAR(255) NOT NULL,
		ordered_at timestamptz NOT NULL DEFAULT (now())
	);
	`
	itemTable := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		item_code VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		quantity INT NOT NULL,
		order_id INT NOT NULL,
		CONSTRAINT items_order_id_fkey
			FOREIGN KEY (order_id)
				REFERENCES orders (id)
					ON DELETE CASCADE
	);
	`

	createTableQueries := fmt.Sprintf("%s %s", orderTable, itemTable)
	_, err := db.Exec(createTableQueries)

	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func InitializeDB() {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, dbPort, dbName)

	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while pinging database: ", err.Error())
	}
	fmt.Println("Successfully connected to database!")
	createRequiredTable()
}

func GetDB() *sql.DB {
	return db
}
