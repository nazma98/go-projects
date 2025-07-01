package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"Apples":  5,
		"Bananas": 7,
	}

	value, ok := myMap["Apples"]
	if ok {
		fmt.Println("Found it ", value)
	} else {
		fmt.Println("Not found")
	}

}
