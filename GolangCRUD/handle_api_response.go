package main

import (
        "fmt"
        "github.com/gorilla/mux"
        "net/http"
//         "encoding/json"
        "log"
        "io"
)



// type Product struct {
// 	ID          int    `json:"id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Category    string `json:"category"` // Modify if API doesn't provide category
// 	Price       float64 `json:"price,omitempty"` // Optional field
// 	// Add other fields as needed based on API response
// }

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Replace with the actual API URL for products
	apiURL := "https://dummyjson.com/products/4"

	// Make HTTP request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is not 200 OK
	if response.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error fetching data: API returned status code %d", response.StatusCode), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header of the response to application/json
	w.Header().Set("Content-Type", "application/json")

	// Copy the response body to the response writer
	_, err = io.Copy(w, response.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error copying response: %v", err), http.StatusInternalServerError)
		return
	}
}



    func main() {
        r := mux.NewRouter()

        r.HandleFunc("/products",ProductsHandler).Methods("GET")
        log.Println("Server listening on :8000")
        log.Fatal(http.ListenAndServe(":8000", r))
    }