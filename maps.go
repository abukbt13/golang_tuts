package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{"abu": "dev"}
	m["alex"] = "software dev"
	fmt.Print(m)
}
