package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Price             float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating            float64  `json:"rating"`
	Stock             int      `json:"stock"`
	Brand             string   `json:"brand"`
	Category          string   `json:"category"`
	Thumbnail         string   `json:"thumbnail"`
	Images            []string `json:"images"`
}

func main() {
	// Make the HTTP GET request to fetch the data.
	resp, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body.
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON data into a Product struct.
	var product Product
	err = json.Unmarshal(responseData, &product)
	if err != nil {
		panic(err)
	}

	// Print the parsed product data.
	fmt.Printf("Product ID: %d\n", product.ID)
	fmt.Printf("Title: %s\n", product.Title)
	fmt.Printf("Description: %s\n", product.Description)
	fmt.Printf("Price: %.2f\n", product.Price)
	fmt.Printf("Discount Percentage: %.2f%%\n", product.DiscountPercentage)
	fmt.Printf("Rating: %.2f\n", product.Rating)
	fmt.Printf("Stock: %d\n", product.Stock)
	fmt.Printf("Brand: %s\n", product.Brand)
	fmt.Printf("Category: %s\n", product.Category)
	fmt.Printf("Thumbnail URL: %s\n", product.Thumbnail)
	fmt.Println("Images:")
	for _, image := range product.Images {
		fmt.Println(image)
	}
}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Price             float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating            float64  `json:"rating"`
	Stock             int      `json:"stock"`
	Brand             string   `json:"brand"`
	Category          string   `json:"category"`
	Thumbnail         string   `json:"thumbnail"`
	Images            []string `json:"images"`
}

func main() {
	// Make the HTTP GET request to fetch the data.
	resp, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body.
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON data into a Product struct.
	var product Product
	err = json.Unmarshal(responseData, &product)
	if err != nil {
		panic(err)
	}

	// Print the parsed product data.
	fmt.Printf("Product ID: %d\n", product.ID)
	fmt.Printf("Title: %s\n", product.Title)
	fmt.Printf("Description: %s\n", product.Description)
	fmt.Printf("Price: %.2f\n", product.Price)
	fmt.Printf("Discount Percentage: %.2f%%\n", product.DiscountPercentage)
	fmt.Printf("Rating: %.2f\n", product.Rating)
	fmt.Printf("Stock: %d\n", product.Stock)
	fmt.Printf("Brand: %s\n", product.Brand)
	fmt.Printf("Category: %s\n", product.Category)
	fmt.Printf("Thumbnail URL: %s\n", product.Thumbnail)
	fmt.Println("Images:")
	for _, image := range product.Images {
		fmt.Println(image)
	}
}
