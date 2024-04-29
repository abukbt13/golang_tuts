package main

import (
	"fmt"
	// "time"
)

func main() {
	// fmt.Printf("showing time")
	// fmt.Printf("The time is", time.Now())

	// b := [5]int{1, 2, 2, 34, 2}
	// slice := b[1:4]
	// fmt.Print(slice)

	newslice := []map[string]interface{}{}

	// Append data to the slice
	newslice = append(newslice, map[string]interface{}{
		"name": "Abuu",
		"age":  20,
		"home": "kaps",
	})
	newslice = append(newslice, map[string]interface{}{
		"name": "Alaa",
		"age":  26,
		"home": "Nairobi",
	})
	fmt.Println(newslice)

}
