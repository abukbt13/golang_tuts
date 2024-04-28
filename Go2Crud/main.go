package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", Isbn: "56789", Title: "Twinkle twinkle", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	movies = append(movies, movie{ID: "1", Isbn: "56789", Title: "Travel of a raindrop", Director: &Director{Firstname: "Sam", Lastname: "Saoong"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Startin server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
