package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = ""
	dbName   = "gocrud_app"
) //database connection

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Parse JSON data from the request body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	ID    int
	Name  string
	Email string
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)

	// Call the GetUser function to fetch the user data from the database
	user, err := GetUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	// Call the GetUser function to fetch the user data from the database
	UpdateUser(db, userID, user.Name, user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User updated successfully")
}
func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err
	}
	return nil
}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	user := DeleteUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User deleted successfully")

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// This function deletes the user with the specified ID from the table. The deleteUserHandler contains, the DeleteUser function which runs the SQL query to delete the user with the provided ID. The delete function(DeleteUser) is as it is written below:

func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// code for adding a record to database
//Invoke-WebRequest -Uri "http://localhost:8090/user" -Method Post -Headers @{"Content-Type"="application/json"} -Body '{"Name":"John Doe","Email":"john@example.com"}'

//request for viewing records in database
// Invoke-WebRequest -Uri "http://localhost:8090/user/1" -Method Delete

// code for updating user
// Invoke-WebRequest -Uri "http://localhost:8090/user/123" -Method Put -Headers @{"Content-Type"="application/json"} -Body '{"Name":"Jane Doe","Email":"jane@example.com"}'

//parameter request for deleting
//  Invoke-WebRequest -Uri "http://localhost:8090/user/1" -Method Delete

//
// go get -u github.com/go-sql-driver/mysql This command downloads and installs the github.com/go-sql-driver/mysql package, which provides the MySQL driver for Go applications.
//go mod init This should generate a go.mod file with the URL you just wrote and the Golang version you are using.
