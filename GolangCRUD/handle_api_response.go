package main

import (
        "fmt"
        "github.com/gorilla/mux"
        "net/http"
//         "encoding/json"
        "log"
        "io/ioutil"
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
	apiURL := "https://dummyjson.com/produts/4"

	// Make HTTP request to the API
	response, err := http.Get(apiURL)
    responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(responseData)
}



    func main() {
        r := mux.NewRouter()

        r.HandleFunc("/products",ProductsHandler).Methods("GET")
        log.Println("Server listening on :8080")
        log.Fatal(http.ListenAndServe(":8080", r))
    }