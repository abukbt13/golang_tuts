package main

import (
	"fmt"
)

func main() {
	var (
		total float64
	)
	var i int

	fmt.Print("Enter a value between 1 and 10 ")
	//want i as the keyboard input
	fmt.Scanln(&i) // Read input from the keyboard and store it in variable i

	switch i {
	case 1:
		{
			total = float64(i) * 3.14
			fmt.Println("The computation value of", i, "is ", total)
		}
	case 2:
		total = float64(i) * 3.14
		fmt.Println("The computation value of", i, "is ", total)
	case 3:
		total = float64(i) * 3.14
		fmt.Println("The computation value of", i, "is ", total)
	default:
		fmt.Println("The value entered is not  computable ")
	}

}
