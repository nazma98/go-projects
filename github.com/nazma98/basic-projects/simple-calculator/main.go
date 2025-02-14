package main

import "fmt"

func getNumbers() (int, int) {
	fmt.Println("Enter first number")
	var firstNumber int
	fmt.Scanln(&firstNumber)
	fmt.Println("Enter second number")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	return firstNumber, secondNumber
}

func main() {
	firstNumber, secondNumber := getNumbers()
	fmt.Println(firstNumber," ", secondNumber)
}