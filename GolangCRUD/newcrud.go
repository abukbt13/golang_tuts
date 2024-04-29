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
	dbPass   = "Susan@2022"
	dbName   = "gocrud_app"
) //database connection

func checkDbconnection(w http.ResponseWriter, r *http.Request) {
    // Attempt to open a connection to the database
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        // If there's an error while opening the connection, return a failure response
        http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Attempt to ping the database to check if the connection is valid
    if err := db.Ping(); err != nil {
        // If the ping fails, return a failure response
        http.Error(w, "Failed to ping the database. Incorrect database name or user credentials.", http.StatusInternalServerError)
        return
    }

    // If there's no error, it means the connection was successful
    // You can return a success response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Successfully connected to the database")
}

func dbConnection() (*sql.DB, error) {
    // Attempt to open a connection to the database
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        return nil, err
    }

    // Ping the database to check if the connection is valid
    if err := db.Ping(); err != nil {
        db.Close() // Close the connection if ping fails
        return nil, err
    }

    return db, nil
}
type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    RegNo string `json:"regno"`
}

func createStudent(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Failed to parse form data", http.StatusBadRequest)
        return
    }

    // Create a new Student instance
    student := Student{
        Name:  r.Form.Get("name"),
        Email: r.Form.Get("email"),
        RegNo: r.Form.Get("regno"),
    }


//     defer db.Close()
 db, err := dbConnection()
    // Prepare SQL statement for inserting student data
    stmt, err := db.Prepare("INSERT INTO students (name, email, regno) VALUES (?, ?, ?)")
    if err != nil {
        http.Error(w, "Failed to prepare SQL statement", http.StatusInternalServerError)
        return
    }
//     defer stmt.Close()

    // Execute SQL statement with student data
    res, err := stmt.Exec(student.Name, student.Email, student.RegNo)
    if err != nil {
        http.Error(w, "Failed to insert student data into database: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Get the auto-generated ID from the database
    studentID, err := res.LastInsertId()
    if err != nil {
        http.Error(w, "Failed to get last insert ID", http.StatusInternalServerError)
        return
    }

    // Set the generated ID in the student object
    student.ID = int(studentID)

    // Return success response
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "Successfully inserted into to the database")
    json.NewEncoder(w).Encode(student)
}
func updateStudent(w http.ResponseWriter, r *http.Request) {
    // Extract student ID from URL path parameters
    vars := mux.Vars(r)
    studentIDStr := vars["id"]
    studentID, err := strconv.Atoi(studentIDStr)
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    // Parse form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Failed to parse form data", http.StatusBadRequest)
        return
    }

    // Create a new Student instance
    student := Student{
        ID:    studentID,
        Name:  r.Form.Get("name"),
        Email: r.Form.Get("email"),
        RegNo: r.Form.Get("regno"),
    }

    // Establish a database connection
    db, err := dbConnection()
    if err != nil {
        http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Prepare SQL statement for updating student data
    stmt, err := db.Prepare("UPDATE students SET name=?, email=?, regno=? WHERE id=?")
    if err != nil {
        http.Error(w, "Failed to prepare SQL statement", http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    // Execute SQL statement with student data
    _, err = stmt.Exec(student.Name, student.Email, student.RegNo, student.ID)
    if err != nil {
        http.Error(w, "Failed to update student data in the database: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Successfully updated student in the database")
    json.NewEncoder(w).Encode(student)
}


func deleteStudent(w http.ResponseWriter, r *http.Request) {
    // Extract student ID from URL path parameters
    vars := mux.Vars(r)
    studentIDStr := vars["id"]
    studentID, err := strconv.Atoi(studentIDStr)
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    // Establish a database connection
    db, err := dbConnection()
    if err != nil {
        http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Prepare SQL statement for deleting student data
    stmt, err := db.Prepare("DELETE FROM students WHERE id=?")
    if err != nil {
        http.Error(w, "Failed to prepare SQL statement", http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    // Execute SQL statement to delete the student record
    _, err = stmt.Exec(studentID)
    if err != nil {
        http.Error(w, "Failed to delete student data from the database: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Successfully deleted student from the database")
}

func showStudents(w http.ResponseWriter, r *http.Request) {
    // Establish a database connection
    db, err := dbConnection()
    if err != nil {
        http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query all students from the database
    rows, err := db.Query("SELECT id, name, email, regno FROM students")
    if err != nil {
        http.Error(w, "Failed to query students from the database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Create a slice to store the retrieved students
    var students []Student

    // Iterate over the rows and populate the students slice
    for rows.Next() {
        var student Student
        if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.RegNo); err != nil {
            http.Error(w, "Failed to scan student data from the database: "+err.Error(), http.StatusInternalServerError)
            return
        }
        students = append(students, student)
    }

    // Check for errors during row iteration
    if err := rows.Err(); err != nil {
        http.Error(w, "Error occurred while iterating over student rows: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Encode the students slice as JSON and return as response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}


func main() {
	r :=mux.NewRouter()

	// Define your HTTP routes using the router
	// r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/testconnection", checkDbconnection).Methods("GET")

	r.HandleFunc("/student", createStudent).Methods("POST")
	r.HandleFunc("/students/show", showStudents).Methods("GET")
	r.HandleFunc("/updateStudent/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/deleteStudent/{id}", deleteStudent).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
