package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["Apples"] = 5
	myMap["Bananas"] = 7

	fmt.Println(myMap)
}
