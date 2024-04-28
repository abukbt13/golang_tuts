package main

import (
	"database/sql"
	// "encoding/json"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "strconv"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "s"
	dbName   = ""
) //database connection

func main() {
	// Create a new router
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database successfully")

	defer db.Close()

	// Check if the connection is actually open
	err = db.Ping()
	if err != nil {
		log.Println("Failed")
		log.Fatal("Failed to ping database:", err)
	}

	// Define your HTTP routes using the router
	// r.HandleFunc("/user", createUserHandler).Methods("POST")
	// r.HandleFunc("/testconnection", dbshowConnection).Methods("GET")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8000")
	// log.Fatal(http.ListenAndServe(":8090", r))
}
