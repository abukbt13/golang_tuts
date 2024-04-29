package main

import "fmt"

func main() {

	var Students = [5]string{"abraham", "kibet", "jacob", "Aooron", "kkqwjk"}
	fmt.Println(Students[2])

	// fmt.Printf("The aray length %v", len(Students))
	scores := []string{"jjkdjks", "njwknjknwdk"}

	// Append elements to the slice
	scores = append(scores, "Gideon", "Victor")

	// Print the updated slice
	fmt.Println("The scores", scores)
	fmt.Println(len(scores))
	// fmt.Println("Welcome to GoLang successs !!!!!!")
	// var name = "Abraham Kibet"
	// var age = "26"
	// var phone = "099999"
	// fmt.Print(name, age, phone)
	// fmt.Print("Success on first code")
	// const (
	// 	PI    = 3.14
	// 	LABEL = "TOYOTA"
	// )
	// print(PI, LABEL)
}
