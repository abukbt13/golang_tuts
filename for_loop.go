package main

import (
	"fmt"
)

func main() {

	var i = 2
	for i < 10 {
		fmt.Print("hello", i, "\n")
		i++
	}
	var MyArray = [5]int{}
	MyArray[4] = 45
	fmt.Print(MyArray)
}
