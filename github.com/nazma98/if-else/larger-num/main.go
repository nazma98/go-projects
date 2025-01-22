package main

import "fmt"

func getNumbers() (int, int){
	var num1 int
	var num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	return num1, num2
}

func main() {
	getNumbers()
}