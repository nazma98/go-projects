package main

import "fmt"

func getNumber() int {
	fmt.Println("Enter the starting number : ")
	var startNumber int
	fmt.Scanln(&startNumber)
	return startNumber
}

func main() {
	startNumber := getNumber()
}