package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/gorilla/mux"
)

// User struct to represent user data
type Client struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// In-memory user storage (replace with your database interaction)
var client []Client

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = ""
	dbName   = "gocrud_app"
)

func connectToDB() (*sql.DB, error) {
	// Replace with your database connection details
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// hashPassword function (replace with your actual hashing implementation using a library like bcrypt)
func authenticateUser(email, password string) bool {
	// Connect to the database
	db, err := connectToDB()
	if err != nil {
		// Handle database connection error
		fmt.Println("Error connecting to the database:", err)
		return false
	}
	defer db.Close()

	// Query the database to check if the email and password match
	var storedPassword string
	err = db.QueryRow("SELECT password FROM clients WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		// Handle query error
		fmt.Println("Error querying the database:", err)
		return false
	}

	// Compare the stored password with the provided password
	if storedPassword == password {
		// Passwords match, authentication successful
		return true
	}

	// Passwords don't match, authentication failed
	return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get form values
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// Authenticate user using email and password
	if authenticateUser(email, password) {
		// If authentication succeeds, respond with success message
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Login successful!")
	} else {
		// If authentication fails, respond with error message
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a Client struct instance
	client := Client{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	// Connect to the database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmt, err := db.Prepare("INSERT INTO clients(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Insert data into the database
	result, err := stmt.Exec(client.Name, client.Email, client.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	clientID, err := result.LastInsertId()
	client.ID = int(clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Data inserted successfully!")

	// Respond with success message and client ID
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Data inserted successfully! Client ID: %d")
	json.NewEncoder(w).Encode(client)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/register", registerHandler).Methods("POST")

	log.Printf("Server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
