package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]int)

	myMap["Books"] = 3
	myMap["Stories"] = 7

	for key, value := range myMap {
		fmt.Println(key, " = ", value)
	}
}
