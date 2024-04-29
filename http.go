package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	// Print the HTTP response status.
	responseData, err := io.ReadAll(resp.Body)
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response data:", responseData)
}
