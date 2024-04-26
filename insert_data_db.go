package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the driver
)

func main() {
	// Database connection details (replace with your own)
	dataSourceName := "root:@tcp(localhost:3306)/students"

	// Connect to the database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Close the connection when the function exits

	// Prepare the INSERT statement
	stmt, err := db.Prepare("INSERT INTO details (name, email) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}

	// Insert data
	name := "John Doe"
	email := "john.doe@example.com"
	_, err = stmt.Exec(name, email)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted successfully!")
}
